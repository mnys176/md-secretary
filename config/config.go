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
		AbstractTemplate       string `toml:"abstract-template"`
		ResourcesTemplate      string `toml:"resources-template"`
		FurtherReadingTemplate string `toml:"further-reading-template"`
	}
	Log struct {
		ContentTemplate string `toml:"content-template"`
	}
	Summary struct {
		SummaryTemplate string `toml:"summary-template"`
		ContentTemplate string `toml:"content-template"`
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
	if custom.Project.AbstractTemplate != "" {
		c.Project.AbstractTemplate = custom.Project.AbstractTemplate
	}
	if custom.Project.ResourcesTemplate != "" {
		c.Project.ResourcesTemplate = custom.Project.ResourcesTemplate
	}
	if custom.Project.FurtherReadingTemplate != "" {
		c.Project.FurtherReadingTemplate = custom.Project.FurtherReadingTemplate
	}
	if custom.Log.ContentTemplate != "" {
		c.Log.ContentTemplate = custom.Log.ContentTemplate
	}
	if custom.Summary.ContentTemplate != "" {
		c.Summary.ContentTemplate = custom.Summary.ContentTemplate
	}
	if custom.Summary.SummaryTemplate != "" {
		c.Summary.SummaryTemplate = custom.Summary.SummaryTemplate
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
	if c.Project.AbstractTemplate == "" {
		return fmt.Errorf(
			"Invalid value for `project.abstract`: %s",
			c.Project.AbstractTemplate,
		)
	}
	if c.Project.ResourcesTemplate == "" {
		return fmt.Errorf(
			"Invalid value for `project.resources`: %s",
			c.Project.ResourcesTemplate,
		)
	}
	if c.Project.FurtherReadingTemplate == "" {
		return fmt.Errorf(
			"Invalid value for `project.further-reading`: %s",
			c.Project.FurtherReadingTemplate,
		)
	}
	if c.Log.ContentTemplate == "" {
		return fmt.Errorf(
			"Invalid value for `log.content`: %s",
			c.Log.ContentTemplate,
		)
	}
	if c.Summary.ContentTemplate == "" {
		return fmt.Errorf(
			"Invalid value for `summary.content`: %s",
			c.Summary.ContentTemplate,
		)
	}
	if c.Summary.SummaryTemplate == "" {
		return fmt.Errorf(
			"Invalid value for `summary.summary`: %s",
			c.Summary.SummaryTemplate,
		)
	}
	return nil
}
