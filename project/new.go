package project

import (
	"strings"
	"time"

	"github.com/mnys176/md-secretary/utils"
)

func NewProject(notebookPath string, title string) *Project {
	title = strings.TrimSpace(title)
	m := NewMarker()
	return &Project{
		Title:       title,
		SystemTitle: utils.Systemify(title),
		Start:       m,
		End:         m,
		Markers:     []Marker{*m},
		Notebook:    notebookPath,
	}
}

func NewMarker() *Marker {
	year, month, day := time.Now().Date()
	return &Marker{time.Date(year, month, day, 0, 0, 0, 0, time.UTC)}
}
