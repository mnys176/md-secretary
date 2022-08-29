package project

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/mnys176/md-secretary/config"
)

type Marker struct {
	Date time.Time
}

func (m Marker) Build(projectPath string) error {
	cfg := config.Defaults()
	var mode string = "January-06"
	if cfg.Notebook.CompactMarkerDirectory {
		mode = "Jan-06"
	}
	markerPath := filepath.Join(projectPath, strings.ToLower(m.Date.Format(mode)))
	logFilePath := filepath.Join(markerPath, "log.md")
	summaryFilePath := filepath.Join(markerPath, "summary.md")

	// create directory
	err := os.Mkdir(markerPath, 0755)
	if err != nil {
		return err
	}

	// create log file
	logFile, err := os.Create(logFilePath)
	defer logFile.Close()
	if err != nil {
		return err
	}

	// create summary file
	summaryFile, err := os.Create(summaryFilePath)
	defer summaryFile.Close()
	if err != nil {
		return err
	}
	return nil
}
