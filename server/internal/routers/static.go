package routers

import (
	f "github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
	"path/filepath"
)

var _ Router[*f.App] = (*StaticRouter)(nil)

type StaticRouter struct {
}

func NewStaticRouter() *StaticRouter {
	return &StaticRouter{}
}

func (s *StaticRouter) Setup(app *f.App) {
	app.Get("/web", static.New(filepath.Join("web", "build")))
	app.Get("/web/*", static.New(filepath.Join("web", "build", "index.html")))

	app.Get("/img/*", static.New(filepath.Join("web", "build", "img")))
	app.Get("/static/*", static.New(filepath.Join("web", "build", "static")))
	app.Get("/manifest.json", static.New(filepath.Join("web", "build", "manifest.json")))
}
