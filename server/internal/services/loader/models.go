package loader

import (
	"github.com/WildEgor/pi-stalker-radio/internal/services/storage"
)

type ILoader interface {
	Load(storage storage.IStorage) error
}
