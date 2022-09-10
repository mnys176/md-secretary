package project

import (
	"time"
	"strings"

	"github.com/mnys176/md-secretary/utils"
)

func NewProject(title string) *Project {
	title = strings.TrimSpace(title)
	return &Project{
		Title:       title,
		SystemTitle: utils.Systemify(title),
		Start:       NewMarker(),
		End:         NewMarker(),
	}
}

func NewMarker() *Marker {
	year, month, day := time.Now().Date()
	return &Marker{time.Date(year, month, day, 0, 0, 0, 0, time.UTC)}
}
