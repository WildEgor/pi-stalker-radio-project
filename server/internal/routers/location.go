package routers

import (
	glh "github.com/WildEgor/pi-stalker-radio/internal/handlers/get_locations"
	sth "github.com/WildEgor/pi-stalker-radio/internal/handlers/get_stations_by_location_id"
	"github.com/gorilla/rpc"
	"net/http"
)

var _ Router[*rpc.Server] = (*LocationRouter)(nil)

type LocationRouter struct {
	gh *glh.GetLocationsHandler
	sh *sth.GetStationsByLocationIDHandler
}

func NewLocationRouter(glh *glh.GetLocationsHandler, sh *sth.GetStationsByLocationIDHandler) *LocationRouter {
	return &LocationRouter{gh: glh, sh: sh}
}

func (r *LocationRouter) GetLocations(req *http.Request, args *glh.GetLocationArgs, reply *glh.GetLocationReply) error {
	return r.gh.Handle(reply)
}

func (r *LocationRouter) GetStations(req *http.Request, args *sth.GetStationsArgs, reply *sth.GetStationsReply) error {
	return r.sh.Handle(args, reply)
}

func (r *LocationRouter) Setup(app *rpc.Server) {
	err := app.RegisterService(r, "LocationService")
	if err != nil {
		return
	}
}
