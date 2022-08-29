package project

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
	"time"

	"github.com/mnys176/md-secretary/config"
	"github.com/mnys176/md-secretary/utils"
)

type (
	Project struct {
		Title       string
		SystemTitle string
		Abstract    string
		Start       time.Time
		End         time.Time
	}
	ProjectTemplateData struct {
		Title          string
		Abstract       string
		Resources      string
		FurtherReading string
	}
)

//go:embed templates/project.tmpl
var projectTemplateTmpl string

func (p Project) String(cfg config.Config) string {
	var startString string = p.Start.Format("Jan '06")
	var endString string = p.End.Format("Jan '06")
	var heading string = p.Title + " (" + startString + " - " + endString + ")"
	if len(p.Abstract) == 0 {
		return heading
	}

	return fmt.Sprintf(
		"%s\n\n%s\n",
		heading,
		utils.ChopString(p.Abstract, cfg.DisplayWidth),
	)
}

func (p Project) Build(path string, cfg *config.Config) error {
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

	// render template and populate with starter content
	projectTemplate := template.Must(template.New("project").Parse(projectTemplateTmpl))
	err = projectTemplate.Execute(projectFile, ProjectTemplateData{
		Title:          p.Title,
		Abstract:       cfg.Project.Abstract,
		Resources:      cfg.Project.Resources,
		FurtherReading: cfg.Project.FurtherReading,
	})
	if err != nil {
		return err
	}

	// build first marker
	m := Marker{Date: p.Start}
	err = m.Build(projectPath, cfg)
	if err != nil {
		return err
	}
	return nil
}
