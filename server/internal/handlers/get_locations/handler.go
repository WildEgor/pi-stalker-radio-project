package handlers

import (
	"github.com/WildEgor/pi-stalker-radio/internal/services/storage"
	"strconv"
)

type GetLocationsHandler struct {
	storage storage.IStorage
}

func NewGetLocationsHandler(storage storage.IStorage) *GetLocationsHandler {
	return &GetLocationsHandler{storage}
}

func (h *GetLocationsHandler) Handle(rpl *GetLocationReply) error {

	locations, err := h.storage.ListAllLocations()
	if err != nil {
		return err
	}

	rpl.Locations = make([]Location, 0, len(locations))
	for _, location := range locations {
		rpl.Locations = append(rpl.Locations, Location{
			Name: location.Name,
			Coordinates: LocationCoordinate{
				Lat:  strconv.Itoa(location.X),
				Long: strconv.Itoa(location.Y),
			},
		})
	}

	return nil
}
