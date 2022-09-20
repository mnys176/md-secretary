package project

import (
	"path/filepath"
	"strings"
	"time"

	"github.com/mnys176/md-secretary/utils"
)

func NewProject(notebookPath string, title string) *Project {
	title = strings.TrimSpace(title)
	systemTitle := utils.Systemify(title)
	projectPath := filepath.Join(notebookPath, systemTitle)

	m := *NewMarker(projectPath)
	markers := []Marker{m}
	return &Project{
		Title:       title,
		SystemTitle: systemTitle,
		Start:       &markers[0],
		End:         &markers[0],
		Markers:     markers,
		Notebook:    notebookPath,
	}
}

func NewMarker(projectPath string) *Marker {
	year, month, day := time.Now().Date()
	return &Marker{
		Project: projectPath,
		Date:    time.Date(year, month, day, 0, 0, 0, 0, time.UTC),
	}
}
