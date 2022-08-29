package project

import (
	_ "embed"
	"os"
	"path/filepath"
	"strings"
	"time"
	"text/template"

	"github.com/mnys176/md-secretary/config"
	"github.com/mnys176/md-secretary/utils"
)

type (
	Marker struct {
		Date time.Time
	}
	LogTemplateData struct {
		Title string
		MarkerDate string
		Date string
		Content string
	}
	SummaryTemplateData struct {
		Title string
		MarkerDate string
		Date string
		Summary string
		Content string
	}
)

//go:embed templates/log.tmpl
var logTemplateTmpl string

//go:embed templates/summary.tmpl
var summaryTemplateTmpl string

func (m Marker) Build(projectPath string, cfg *config.Config) error {
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

	// render template and populate log file with starter content
	logTemplate := template.Must(template.New("log").Parse(logTemplateTmpl))
	err = logTemplate.Execute(logFile, LogTemplateData{
		Title: utils.Desystemify(filepath.Base(projectPath)),
		MarkerDate: m.Date.Format("January, 2006"),
		Date: m.Date.Format("Monday, 01/02"),
		Content: cfg.Log.Content,
	})
	if err != nil {
		return err
	}

	// create summary file
	summaryFile, err := os.Create(summaryFilePath)
	defer summaryFile.Close()
	if err != nil {
		return err
	}

	// render template and populate summary file with starter content
	summaryTemplate := template.Must(template.New("summary").Parse(summaryTemplateTmpl))
	err = summaryTemplate.Execute(summaryFile, SummaryTemplateData{
		Title: utils.Desystemify(filepath.Base(projectPath)),
		MarkerDate: m.Date.Format("January, 2006"),
		Date: m.Date.Format("Monday, 01/02"),
		Summary: cfg.Summary.Summary,
		Content: cfg.Summary.Content,
	})
	if err != nil {
		return err
	}
	return nil
}
