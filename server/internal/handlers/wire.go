package handlers

import (
	glh "github.com/WildEgor/pi-stalker-radio/internal/handlers/get_locations"
	gsh "github.com/WildEgor/pi-stalker-radio/internal/handlers/get_stations_by_location_id"
	hch "github.com/WildEgor/pi-stalker-radio/internal/handlers/health_check"
	"github.com/google/wire"
)

// Set contains http/amqp/etc hch (acts like facades)
var Set = wire.NewSet(
	hch.NewHealthCheckHandler,
	glh.NewGetLocationsHandler,
	gsh.NewGetStationsByLocationIDHandler,
)
