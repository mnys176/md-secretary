package project

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/mnys176/md-secretary/config"
	"github.com/mnys176/md-secretary/utils"
)

type (
	Project struct {
		Title       string
		SystemTitle string
		Abstract    string
		Start       *Marker
		End         *Marker
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

func (p Project) String(cfg *config.Config) string {
	var (
		titleContentWidth    int = cfg.DisplayWidth - 24
		markerContentWidth   int = 17
		abstractContentWidth int = cfg.DisplayWidth - 4
	)
	var start string = p.Start.Date.Format("Jan '06")
	var end string = p.End.Date.Format("Jan '06")
	var builder strings.Builder

	// build border then reset
	builder.WriteString("+")
	builder.WriteString(strings.Repeat("-", titleContentWidth+2))
	builder.WriteString("+")
	builder.WriteString(strings.Repeat("-", markerContentWidth+2))
	builder.WriteString("+\n")
	border := builder.String()
	builder.Reset()
	builder.WriteString(border)

	// ensure title fits in content area
	titleLines := strings.Split(utils.ChopString(p.Title, titleContentWidth), "\n")
	for i, line := range titleLines {
		builder.WriteString("| ")
		builder.WriteString(line)
		builder.WriteString(strings.Repeat(" ", titleContentWidth-len(line)))
		builder.WriteString(" | ")
		if i == 0 {
			builder.WriteString(start)
			builder.WriteString(" - ")
			builder.WriteString(end)
		} else {
			builder.WriteString(strings.Repeat(" ", markerContentWidth))
		}
		builder.WriteString(" |\n")
	}
	builder.WriteString(border)

	if len(p.Abstract) == 0 {
		return strings.TrimSuffix(builder.String(), "\n")
	}

	// ensure abstract fits in content area
	abstractLines := strings.Split(utils.ChopString(p.Abstract, abstractContentWidth), "\n")
	for _, line := range abstractLines {
		line = strings.TrimSuffix(line, " ")
		builder.WriteString("| ")
		builder.WriteString(line)
		builder.WriteString(strings.Repeat(" ", abstractContentWidth-len(line)))
		builder.WriteString(" |\n")
	}
	border = strings.ReplaceAll(border, "-+-", "---")
	builder.WriteString(border)
	return strings.TrimSuffix(builder.String(), "\n")
}

func (p Project) Build(notebookPath string, cfg *config.Config) error {
	projectPath := filepath.Join(notebookPath, p.SystemTitle)
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
	err = p.Start.Build(projectPath, cfg)
	if err != nil {
		return err
	}
	return nil
}

func (p Project) Append(notebookPath string, cfg *config.Config) error {
	projectPath := filepath.Join(notebookPath, p.SystemTitle)
	m := NewMarker()
	err := m.Build(projectPath, cfg)
	if err != nil {
		return err
	}
	return nil
}

func (p Project) TearDown(notebookPath string, force bool, cfg *config.Config) error {
	projectPath := filepath.Join(notebookPath, p.SystemTitle)
	dialog := fmt.Sprintf("Scrapping the project `%s` will permanently"+
		" delete all associated files, including media,"+
		" from existence.", p.Title)
	if force || utils.Confirm(utils.ChopString(dialog, cfg.DisplayWidth)) {
		if !force {
			fmt.Println("\nScrapping project...")
		}
		err := os.RemoveAll(projectPath)
		if err != nil {
			return err
		}
	}
	return nil
}
