package project

import (
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/mnys176/md-secretary/utils"
)

func Load(notebookPath string, title string) (*Project, error) {
	title = strings.TrimSpace(title)
	p := Project{
		Title:       utils.Desystemify(title),
		SystemTitle: utils.Systemify(title),
		Notebook:    notebookPath,
	}

	// ensure project exists
	projectPath := filepath.Join(p.Notebook, p.SystemTitle)
	_, err := os.Stat(projectPath)
	if err != nil {
		return nil, err
	}

	files, err := os.ReadDir(projectPath)
	if err != nil {
		return nil, err
	}

	// aggregate marker directories
	for _, f := range files {
		matched, _ := regexp.MatchString(`[a-z]{3}-\d{2}`, f.Name())
		if matched {
			t, _ := time.Parse("Jan-06", f.Name())
			p.Markers = append(p.Markers, Marker{
				Project: projectPath,
				Date:    t,
			})
		}
	}

	// find the start and end markers in the project
	sort.Slice(p.Markers, func(i int, j int) bool {
		return p.Markers[i].Date.Before(p.Markers[j].Date)
	})
	p.Start = &p.Markers[0]
	p.End = &p.Markers[len(p.Markers)-1]

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
