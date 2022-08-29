package ingest

import (
	"fmt"
	"path/filepath"

	"github.com/mnys176/md-secretary/config"
)

func Handle(e *Ingest) error {
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
	fmt.Println(e)
	return nil
}
