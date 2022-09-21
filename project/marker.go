package project

import (
	_ "embed"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"
	// "regexp"

	"github.com/mnys176/md-secretary/config"
	"github.com/mnys176/md-secretary/utils"
)

type (
	Marker struct {
		Project string
		Date    time.Time
		Compact bool
	}
	LogTemplateData struct {
		Title      string
		MarkerDate string
		Date       string
		Content    string
	}
	SummaryTemplateData struct {
		Title      string
		MarkerDate string
		Date       string
		Summary    string
		Content    string
	}
)

//go:embed templates/log.tmpl
var logTemplateTmpl string

//go:embed templates/summary.tmpl
var summaryTemplateTmpl string

func (m Marker) Build(cfg *config.Config) error {
	var mode string = "January-2006"
	if m.Compact {
		mode = "Jan-06"
	}
	markerPath := filepath.Join(m.Project, strings.ToLower(m.Date.Format(mode)))
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
		Title:      utils.Desystemify(filepath.Base(m.Project)),
		MarkerDate: m.Date.Format("January, 2006"),
		Date:       m.Date.Format("Monday, 01/02"),
		Content:    cfg.Log.ContentTemplate,
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
		Title:      utils.Desystemify(filepath.Base(m.Project)),
		MarkerDate: m.Date.Format("January, 2006"),
		Date:       m.Date.Format("Monday, 01/02"),
		Summary:    cfg.Summary.SummaryTemplate,
		Content:    cfg.Summary.ContentTemplate,
	})
	if err != nil {
		return err
	}
	return nil
}

func (m Marker) MarshalJSON() ([]byte, error) {
	type MarkerJson struct {
		Date        int64 `json:"date"`
		LogFile     File  `json:"logFile"`
		SummaryFile File  `json:"summaryFile"`
	}

	var mode string = "January-2006"
	if m.Compact {
		mode = "Jan-06"
	}
	markerPath := filepath.Join(m.Project, strings.ToLower(m.Date.Format(mode)))
	logFilePath := filepath.Join(markerPath, "log.md")
	summaryFilePath := filepath.Join(markerPath, "summary.md")

	// compress log file
	logFileBytes, err := os.ReadFile(logFilePath)
	if err != nil {
		return nil, err
	}
	logFileCompressed, err := utils.CompressEncode(logFileBytes)
	if err != nil {
		return nil, err
	}

	// compress summary file
	summaryFileBytes, err := os.ReadFile(summaryFilePath)
	if err != nil {
		return nil, err
	}
	summaryFileCompressed, err := utils.CompressEncode(summaryFileBytes)
	if err != nil {
		return nil, err
	}

	return json.Marshal(MarkerJson{
		Date:        m.Date.Unix(),
		LogFile:     File{"log.md", logFileCompressed},
		SummaryFile: File{"summary.md", summaryFileCompressed},
	})
}
