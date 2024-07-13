package storage

import (
	"github.com/WildEgor/pi-stalker-radio/internal/domain"
	"sync"
)

var _ IStorage = (*MemoryStorage)(nil)

type MemoryStorage struct {
	locations *sync.Map
	radios    *sync.Map

	locationsCount int

	radioPerPage int
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		locations:    new(sync.Map),
		radios:       new(sync.Map),
		radioPerPage: 10,
	}
}

func (s *MemoryStorage) Save(location *domain.Location) error {
	s.locations.Store(location.ID, location)
	s.locationsCount++

	for _, feature := range location.Features {

		switch feature.PropertiesUnpacked.(type) {
		case *[]domain.RadioStation:
			var radios []*domain.RadioStation

			for _, radio := range *feature.PropertiesUnpacked.(*[]domain.RadioStation) {
				radios = append(radios, &radio)
			}

			s.radios.Store(location.ID, radios)
		}
	}

	return nil
}

func (s *MemoryStorage) ListAllLocations() ([]*domain.Location, error) {
	list := make([]*domain.Location, 0, s.locationsCount)

	s.locations.Range(func(key, value any) bool {
		list = append(list, value.(*domain.Location))
		return true
	})

	return list, nil
}

func (s *MemoryStorage) FindStationsByLocationID(id string) ([]*domain.RadioStation, error) {
	v, ok := s.radios.Load(id)
	if !ok {
		return nil, ErrEmptyPlaylist
	}

	return v.([]*domain.RadioStation), nil
}

func (s *MemoryStorage) Finalize() error {
	return nil
}

func (s *MemoryStorage) Close() {
	s.locations = &sync.Map{}
	s.radios = &sync.Map{}
}
