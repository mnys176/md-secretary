package project

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/mnys176/md-secretary/utils"
)

func Load(notebookPath string, title string) (*Project, error) {
	title = strings.TrimSpace(title)
	p := Project{
		Title:       title,
		SystemTitle: utils.Systemify(title),
	}

	// ensure project exists
	projectPath := filepath.Join(notebookPath, p.SystemTitle)
	_, err := os.Stat(projectPath)
	if err != nil {
		return nil, err
	}

	files, err := os.ReadDir(projectPath)
	if err != nil {
		return nil, err
	}

	// aggregate marker directory names to determine project lifespan
	markerDirectories := []string{}
	for _, f := range files {
		matched, _ := regexp.MatchString(`[a-z]{3}-\d{2}`, f.Name())
		if matched {
			markerDirectories = append(markerDirectories, f.Name())
		}
	}

	// find the oldest and newest markers in the project
	newest := time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)
	oldest := time.Date(9999, time.December, 1, 0, 0, 0, 0, time.UTC)
	for _, m := range markerDirectories {
		t, _ := time.Parse("Jan-06", m)
		if t.Before(oldest) {
			oldest = t
		}
		if t.After(newest) {
			newest = t
		}
	}

	// skip generators and assign dates directly
	p.Start = &Marker{oldest}
	p.End = &Marker{newest}

	projectFilePath := filepath.Join(projectPath, p.SystemTitle+".md")
	projectFileBytes, err := os.ReadFile(projectFilePath)
	if err != nil {
		return nil, err
	}

	// project abstract is between the title and resources headings
	titleHeadingRegex := regexp.MustCompile(`# [\w ]*\n\n`)
	resourcesHeadingRegex := regexp.MustCompile(`\n\n#{2} Resources`)
	abstractStart := titleHeadingRegex.FindIndex(projectFileBytes)[1]
	abstractEnd := resourcesHeadingRegex.FindIndex(projectFileBytes)[0]

	// abstract is not present if this is false
	if abstractStart < abstractEnd {
		p.Abstract = strings.TrimSpace(string(projectFileBytes[abstractStart:abstractEnd]))
	}

	return &p, nil
}
