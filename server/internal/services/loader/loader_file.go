package loader

import (
	"encoding/json"
	"github.com/WildEgor/pi-stalker-radio/internal/configs"
	"github.com/WildEgor/pi-stalker-radio/internal/domain"
	"github.com/WildEgor/pi-stalker-radio/internal/services/storage"
	"os"
	"path/filepath"
)

var _ ILoader = (*FileLoader)(nil)

type FileLoader struct {
	path string
}

func NewFileLoader(ac *configs.AppConfig) *FileLoader {
	return &FileLoader{
		path: ac.AssetsPath,
	}
}

func (f *FileLoader) Load(storage storage.IStorage) error {
	if err := filepath.Walk(f.path, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !f.IsDir() {
			if f.Name() == "locations.json" {
				var locs []domain.Location

				fileContents, err := os.ReadFile(path)
				if err != nil {
					return err
				}

				if err := json.Unmarshal(fileContents, &locs); err != nil {
					return err
				}

				for _, loc := range locs {
					err := storage.Save(&loc)
					if err != nil {
						return err
					}
				}
			}
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}
