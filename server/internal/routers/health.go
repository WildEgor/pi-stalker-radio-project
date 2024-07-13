package routers

import (
	hch "github.com/WildEgor/pi-stalker-radio/internal/handlers/health_check"
	"github.com/gorilla/rpc"
	"net/http"
)

var _ Router[*rpc.Server] = (*HealthRouter)(nil)

type HealthRouter struct {
	hh *hch.HealthCheckHandler
}

func NewHealthRouter(hh *hch.HealthCheckHandler) *HealthRouter {
	return &HealthRouter{hh: hh}
}

// Check example RPC method. In this method convert incoming data to handler input and vice versa
func (r *HealthRouter) Check(req *http.Request, args *hch.HealthArgs, reply *hch.HealthReply) error {
	// mapping here
	return r.hh.Handle()
}

func (r *HealthRouter) Setup(app *rpc.Server) {
	err := app.RegisterService(r, "HealthService")
	if err != nil {
		return
	}
}
