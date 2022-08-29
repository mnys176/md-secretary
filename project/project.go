package project

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/mnys176/md-secretary/config"
	"github.com/mnys176/md-secretary/utils"
)

type Project struct {
	Title       string
	SystemTitle string
	Abstract    string
	Start       time.Time
	End         time.Time
}

func (p Project) String() string {
	var startString string = p.Start.Format("Jan '06")
	var endString string = p.End.Format("Jan '06")
	var heading string = p.Title + " (" + startString + " - " + endString + ")"
	if len(p.Abstract) == 0 {
		return heading
	}

	return fmt.Sprintf(
		"%s\n\n%s\n",
		heading,
		utils.ChopString(p.Abstract, config.Defaults().DisplayWidth),
	)
}

func (p Project) Build(path string) error {
	projectPath := filepath.Join(path, p.SystemTitle)
	projectFilePath := filepath.Join(projectPath, p.SystemTitle+".md")
	mediaPath := filepath.Join(projectPath, "media")

	// create directories
	err := os.Mkdir(projectPath, 0755)
	if err != nil {
		return err
	}
	err = os.Mkdir(mediaPath, 0755)
	if err != nil {
		return err
	}

	// create project file
	projectFile, err := os.Create(projectFilePath)
	defer projectFile.Close()
	if err != nil {
		return err
	}

	// build first marker
	m := Marker{Date: p.Start}
	err = m.Build(projectPath)
	if err != nil {
		return err
	}
	return nil
}
