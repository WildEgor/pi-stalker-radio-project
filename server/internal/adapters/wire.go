package adapters

import (
	"github.com/WildEgor/pi-stalker-radio/internal/adapters/mplayer"
	"github.com/WildEgor/pi-stalker-radio/internal/adapters/pi"
	"github.com/google/wire"
)

// Set contains "adapters" to 3th party systems
var Set = wire.NewSet(
	mplayer.NewMPlayer,
	pi.NewPIWrapper,
)
