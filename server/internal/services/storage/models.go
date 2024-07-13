package storage

import (
	"errors"
	"github.com/WildEgor/pi-stalker-radio/internal/domain"
)

var (
	ErrEmptyPlaylist = errors.New("empty playlist")
)

type IStorage interface {
	Save(location *domain.Location) error
	ListAllLocations() ([]*domain.Location, error)
	FindStationsByLocationID(id string) ([]*domain.RadioStation, error)
	Finalize() error
	Close()
}
