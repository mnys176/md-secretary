package config

import (
	_ "embed"
	"fmt"
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
		CompactMarkerDirectory string `toml:"compact-marker-directory"`
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

	err = c.validate()
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (c *Config) Customize(path string) error {
	custom := Config{}
	absPath, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	_, err = toml.DecodeFile(absPath, &custom)
	if err != nil {
		return err
	}

	if custom.DisplayWidth > 0 {
		c.DisplayWidth = custom.DisplayWidth
	}
	if custom.Notebook.Path != "" {
		c.Notebook.Path = custom.Notebook.Path
	}
	if custom.Notebook.CompactMarkerDirectory != "" {
		c.Notebook.CompactMarkerDirectory = custom.Notebook.CompactMarkerDirectory
	}
	if custom.Compression.Path != "" {
		c.Compression.Path = custom.Compression.Path
	}
	if custom.Compression.JsonTitle != "" {
		c.Compression.JsonTitle = custom.Compression.JsonTitle
	}
	if custom.Project.Abstract != "" {
		c.Project.Abstract = custom.Project.Abstract
	}
	if custom.Project.Resources != "" {
		c.Project.Resources = custom.Project.Resources
	}
	if custom.Project.FurtherReading != "" {
		c.Project.FurtherReading = custom.Project.FurtherReading
	}
	if custom.Log.Content != "" {
		c.Log.Content = custom.Log.Content
	}
	if custom.Summary.Content != "" {
		c.Summary.Content = custom.Summary.Content
	}
	if custom.Summary.Summary != "" {
		c.Summary.Summary = custom.Summary.Summary
	}

	err = c.validate()
	if err != nil {
		return err
	}
	return nil
}

func (c Config) validate() error {
	if c.DisplayWidth < 56 || c.DisplayWidth > 256 {
		return fmt.Errorf(
			"Invalid value for `display-width`: `%d`",
			c.DisplayWidth,
		)
	}
	if c.Notebook.Path == "" {
		return fmt.Errorf(
			"Invalid value for `notebook.path`: `%s`",
			c.Notebook.Path,
		)
	}

	if !map[string]bool{
		"compact":     true,
		"comfortable": true,
	}[c.Notebook.CompactMarkerDirectory] {
		return fmt.Errorf(
			"Invalid value for `notebook.compact-marker-directory`: %s",
			c.Notebook.CompactMarkerDirectory,
		)
	}

	if c.Compression.Path == "" {
		return fmt.Errorf(
			"Invalid value for `compression.path`: %s",
			c.Compression.Path,
		)
	}
	if c.Compression.JsonTitle == "" {
		return fmt.Errorf(
			"Invalid value for `compression.json-title`: %s",
			c.Compression.JsonTitle,
		)
	}
	if c.Project.Abstract == "" {
		return fmt.Errorf(
			"Invalid value for `project.abstract`: %s",
			c.Project.Abstract,
		)
	}
	if c.Project.Resources == "" {
		return fmt.Errorf(
			"Invalid value for `project.resources`: %s",
			c.Project.Resources,
		)
	}
	if c.Project.FurtherReading == "" {
		return fmt.Errorf(
			"Invalid value for `project.further-reading`: %s",
			c.Project.FurtherReading,
		)
	}
	if c.Log.Content == "" {
		return fmt.Errorf(
			"Invalid value for `log.content`: %s",
			c.Log.Content,
		)
	}
	if c.Summary.Content == "" {
		return fmt.Errorf(
			"Invalid value for `summary.content`: %s",
			c.Summary.Content,
		)
	}
	if c.Summary.Summary == "" {
		return fmt.Errorf(
			"Invalid value for `summary.summary`: %s",
			c.Summary.Summary,
		)
	}
	return nil
}
