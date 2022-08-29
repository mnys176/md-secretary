package config

import (
	_ "embed"

	"github.com/BurntSushi/toml"
)

type Config struct {
	DisplayWidth int `toml:"display-width"`
	Project struct {
		Abstract string `toml:"abstract"`
		Resources string `toml:"resources"`
		FurtherReading string `toml:"further-reading"`
	}
	Log struct {
		Content string `toml:"content"`
	}
	Summary struct {
		Summary string `toml:"summary"`
		Content string `toml:"content"`
	}
	Notebook     struct {
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

func Defaults() *Config {
	c := Config{}
	_, err := toml.Decode(defaultConfigToml, &c)
	if err != nil {
		panic(err)
	}
	return &c
}
