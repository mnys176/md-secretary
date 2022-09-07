package project

import (
	"time"

	"github.com/mnys176/md-secretary/utils"
)

func New(title string) *Project {
	year, month, day := time.Now().Date()
	return &Project{
		Title:       title,
		SystemTitle: utils.Systemify(title),
		Start:       Marker{time.Date(year, month, day, 0, 0, 0, 0, time.UTC)},
		End:         Marker{time.Date(year, month, day, 0, 0, 0, 0, time.UTC)},
	}
}
