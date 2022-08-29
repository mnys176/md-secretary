package config

import (
	_ "embed"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type Config struct {
	DisplayWidth int `toml:"display-width"`
	Project      struct {
		Abstract       string `toml:"abstract"`
		Resources      string `toml:"resources"`
		FurtherReading string `toml:"further-reading"`
	}
	Log struct {
		Content string `toml:"content"`
	}
	Summary struct {
		Summary string `toml:"summary"`
		Content string `toml:"content"`
	}
	Notebook struct {
		Path                   string `toml:"path"`
		CompactMarkerDirectory bool   `toml:"compact-marker-directory"`
	}
	Compression struct {
		Path      string `toml:"path"`
		JsonTitle string `toml:"json-title"`
	}
}

//go:embed default.toml
var defaultConfigToml string

func Defaults() (*Config, error) {
	c := Config{}
	_, err := toml.Decode(defaultConfigToml, &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func Custom(path string) (*Config, error) {
	c := Config{}
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	_, err = toml.DecodeFile(absPath, &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
