package domain

import (
	"encoding/json"
	"errors"
)

var (
	ErrUnknownType = errors.New("unknown type")
	ErrUnmarshall  = errors.New("error unmarshal properties")
)

type FeatureJSON struct {
	Type               string          `json:"type"` // "playlist"
	PropertiesRaw      json.RawMessage `json:"properties"`
	PropertiesUnpacked any             `json:"-"`
}

type Feature struct {
	Type               string
	PropertiesRaw      json.RawMessage
	PropertiesUnpacked any
}

func (f *Feature) IsPlaylist() bool {
	return f.Type == "playlist"
}

func (f *Feature) UnmarshalJSON(data []byte) error {
	var fd FeatureJSON

	if err := json.Unmarshal(data, &fd); err != nil {
		return err
	}

	switch fd.Type {
	case "playlist":
		fd.PropertiesUnpacked = &[]RadioStation{}
	default:
		return ErrUnknownType
	}

	if err := json.Unmarshal(fd.PropertiesRaw, fd.PropertiesUnpacked); err != nil {
		return ErrUnmarshall
	}

	f.Type = fd.Type
	f.PropertiesUnpacked = fd.PropertiesUnpacked

	return nil
}
