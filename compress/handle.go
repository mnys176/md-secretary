package compress

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mnys176/md-secretary/config"
	"github.com/mnys176/md-secretary/project"
)

func Handle(e *Compress) error {
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

		// ensure path to custom configuration exists
		_, err = os.Stat(e.Config)
		if err != nil {
			return err
		}

		// load custom configuration
		err = cfg.Customize(e.Config)
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

	// check if output was not provided
	if len(e.Output) == 0 {
		e.Output, err = filepath.Abs(cfg.Compression.Path)
		if err != nil {
			return err
		}
	}
	// fmt.Println(e)

	p, err := project.Load(e.Path, e.ProjectTitle)
	if err != nil {
		return err
	}
	err = p.Export(e.Path, e.Output, e.Transfer, cfg)
	if err != nil {
		return err
	}
	return nil
}
