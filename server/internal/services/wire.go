package services

import (
	"github.com/WildEgor/pi-stalker-radio/internal/adapters"
	"github.com/WildEgor/pi-stalker-radio/internal/services/loader"
	"github.com/WildEgor/pi-stalker-radio/internal/services/storage"
	"github.com/google/wire"
)

// Set contains services
var Set = wire.NewSet(
	adapters.Set,
	loader.NewFileLoader,
	wire.Bind(new(loader.ILoader), new(*loader.FileLoader)),
	storage.NewMemoryStorage,
	wire.Bind(new(storage.IStorage), new(*storage.MemoryStorage)),
)
