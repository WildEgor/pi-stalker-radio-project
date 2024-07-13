// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package pkg

import (
	"github.com/WildEgor/pi-stalker-radio/internal/configs"
	handlers2 "github.com/WildEgor/pi-stalker-radio/internal/handlers/get_locations"
	handlers3 "github.com/WildEgor/pi-stalker-radio/internal/handlers/get_stations_by_location_id"
	"github.com/WildEgor/pi-stalker-radio/internal/handlers/health_check"
	"github.com/WildEgor/pi-stalker-radio/internal/routers"
	"github.com/WildEgor/pi-stalker-radio/internal/services/loader"
	"github.com/WildEgor/pi-stalker-radio/internal/services/storage"
	"github.com/google/wire"
)

// Injectors from server.go:

// NewServer
func NewServer() (*Server, error) {
	configurator := configs.NewConfigurator()
	appConfig := configs.NewAppConfig()
	loggerConfig := configs.NewLoggerConfig()
	healthCheckHandler := handlers.NewHealthCheckHandler()
	healthRouter := routers.NewHealthRouter(healthCheckHandler)
	staticRouter := routers.NewStaticRouter()
	memoryStorage := storage.NewMemoryStorage()
	getLocationsHandler := handlers2.NewGetLocationsHandler(memoryStorage)
	getStationsByLocationIDHandler := handlers3.NewGetStationsByLocationIDHandler(memoryStorage)
	locationRouter := routers.NewLocationRouter(getLocationsHandler, getStationsByLocationIDHandler)
	fileLoader := loader.NewFileLoader(appConfig)
	server := NewApp(configurator, appConfig, loggerConfig, healthRouter, staticRouter, locationRouter, fileLoader, memoryStorage)
	return server, nil
}

// server.go:

// ServerSet
var ServerSet = wire.NewSet(AppSet)
