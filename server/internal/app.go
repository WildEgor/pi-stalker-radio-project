package pkg

import (
	"context"
	"fmt"
	"github.com/WildEgor/pi-stalker-radio/internal/configs"
	"github.com/WildEgor/pi-stalker-radio/internal/routers"
	"github.com/WildEgor/pi-stalker-radio/internal/services"
	"github.com/WildEgor/pi-stalker-radio/internal/services/loader"
	"github.com/WildEgor/pi-stalker-radio/internal/services/storage"
	f "github.com/gofiber/fiber/v3"
	fcors "github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/google/wire"
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"github.com/rs/cors"
	"log/slog"
	"net/http"
	"os"
	"time"
)

// AppSet link main app deps
var AppSet = wire.NewSet(
	NewApp,
	configs.Set,
	routers.Set,
	services.Set,
)

// Server represents the main server configuration.
type Server struct {
	ac  *configs.AppConfig
	app *f.App
	rpc *http.Server
	r   http.Handler
}

// Run start service with deps
func (srv *Server) Run(ctx context.Context) {
	go func() {
		slog.Info("watch config")
		select {
		default:
			if err := srv.app.Listen(fmt.Sprintf(":%s", srv.ac.HTTPPort), f.ListenConfig{
				DisableStartupMessage: false,
				EnablePrintRoutes:     false,
				OnShutdownSuccess: func() {
					slog.Debug("success shutdown service")
				},
			}); err != nil {
				slog.Error("unable to start server")
				os.Exit(1)
			}
		case <-ctx.Done():
			return
		}
	}()

	server := &http.Server{
		Handler:      srv.r,
		Addr:         fmt.Sprintf(":%s", srv.ac.RPCPort),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	slog.Info("listen rpc")

	if err := server.ListenAndServe(); err != nil {
		slog.Error("fail start server", slog.Any("err", err))
		os.Exit(1)
	}
}

// Shutdown graceful shutdown
func (srv *Server) Shutdown(ctx context.Context) {
	if err := srv.app.Shutdown(); err != nil {
		slog.Error("unable to shutdown server")
	}

	err := srv.rpc.Close()
	if err != nil {
		slog.Error("fail stop server", slog.Any("err", err))
		return
	}
}

// NewApp init app
func NewApp(
	c *configs.Configurator,
	ac *configs.AppConfig,
	lc *configs.LoggerConfig, // init logger
	hr *routers.HealthRouter,
	sr *routers.StaticRouter,
	lr *routers.LocationRouter,

	ls loader.ILoader,
	ss storage.IStorage,
) *Server {
	err := ls.Load(ss)
	if err != nil {
		panic(err)
	}

	app := f.New(f.Config{
		AppName:      ac.Name,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  30 * time.Second,
	})

	app.Use(fcors.New(fcors.Config{
		AllowMethods: []string{"GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS"},
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin"},
	}))

	app.Use(recover.New())

	sr.Setup(app)

	s := rpc.NewServer()
	s.RegisterCodec(json.NewCodec(), "application/json")

	hr.Setup(s)
	lr.Setup(s)

	r := mux.NewRouter()

	ds := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://foo.com", "http://foo.com:8080"},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})

	handler := ds.Handler(r)

	r.Handle("/rpc", s)

	return &Server{
		app: app,
		ac:  ac,
		r:   handler,
	}
}
