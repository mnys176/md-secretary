package create

import (
	"fmt"
	"time"
	"path/filepath"

	"github.com/mnys176/md-secretary/config"
	"github.com/mnys176/md-secretary/project"
	"github.com/mnys176/md-secretary/utils"
)

func Handle(e *Create) error {
	if e.Help {
		fmt.Println(Usage)
		return nil
	}

	// load default configuration
	cfg, err := config.Defaults()
	if err != nil {
		return err
	}

	// check if custom configuration provided
	if len(e.Config) > 0 {
		e.Config, err = filepath.Abs(e.Config)
		if err != nil {
			return err
		}

		cfg, err = config.Custom(e.Config)
		if err != nil {
			return err
		}
	}

	// check if path was not provided
	if len(e.Path) == 0 {
		e.Path, err = filepath.Abs(cfg.Notebook.Path)
		if err != nil {
			return err
		}
	}

	year, month, day := time.Now().Date()
	p := project.Project{
		Title:       e.ProjectTitle,
		SystemTitle: utils.Systemify(e.ProjectTitle),
		Start:       time.Date(year, month, day, 0, 0, 0, 0, time.UTC),
		End:         time.Date(year, month, day, 0, 0, 0, 0, time.UTC),
	}

	err = p.Build(e.Path, cfg)
	if err != nil {
		return err
	}
	return nil
}
