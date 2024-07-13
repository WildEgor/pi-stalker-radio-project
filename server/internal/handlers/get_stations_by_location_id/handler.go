package handlers

import (
	"github.com/WildEgor/pi-stalker-radio/internal/services/storage"
)

type GetStationsByLocationIDHandler struct {
	storage storage.IStorage
}

func NewGetStationsByLocationIDHandler(storage storage.IStorage) *GetStationsByLocationIDHandler {
	return &GetStationsByLocationIDHandler{storage}
}

func (h *GetStationsByLocationIDHandler) Handle(args *GetStationsArgs, rpl *GetStationsReply) error {
	stations, err := h.storage.FindStationsByLocationID(args.LocationID)
	if err != nil {
		return err
	}

	rpl.Stations = make([]StationReply, 0, len(stations))
	for _, station := range stations {
		rpl.Stations = append(rpl.Stations, StationReply{
			Name: station.Name,
			URL:  station.URL,
		})
	}

	return nil
}
