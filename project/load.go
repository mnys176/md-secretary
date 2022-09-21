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
		compact, _ := regexp.MatchString(`^[a-z]{3}-\d{2}$`, f.Name())
		comfortable, _ := regexp.MatchString(`^[a-z]{3,9}-\d{4}$`, f.Name())

		var mode string
		if compact {
			mode = "Jan-06"
		} else if comfortable {
			mode = "January-2006"
		}
		if compact || comfortable {
			t, _ := time.Parse(mode, f.Name())
			p.Markers = append(p.Markers, Marker{
				Project: projectPath,
				Date:    t,
				Compact: compact,
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
	resourcesHeadingRegex := regexp.MustCompile(`\n\n#{2} Resources\n\n`)
	furtherReadingHeadingRegex := regexp.MustCompile(`\n\n#{3} Further Reading\n\n`)

	titleHeadingIndeces := titleHeadingRegex.FindIndex(projectFileBytes)
	resourcesHeadingIndeces := resourcesHeadingRegex.FindIndex(projectFileBytes)
	furtherReadingHeadingIndeces := furtherReadingHeadingRegex.FindIndex(projectFileBytes)

	abstractStart := titleHeadingIndeces[1]
	abstractEnd := resourcesHeadingIndeces[0]
	resourcesStart := resourcesHeadingIndeces[1]
	resourcesEnd := furtherReadingHeadingIndeces[0]
	furtherReadingStart := furtherReadingHeadingIndeces[1]

	// these are necessary checks because they exist in the middle
	if abstractStart < abstractEnd {
		p.Abstract = strings.TrimSpace(string(projectFileBytes[abstractStart:abstractEnd]))
	}
	resourceRegex := regexp.MustCompile(`^\* (\[([\w ]+)\]\((.+?)\))$`)
	if resourcesStart < resourcesEnd {
		resources := strings.TrimSpace(string(projectFileBytes[resourcesStart:resourcesEnd]))
		for _, r := range strings.Split(resources, "\n") {
			if !resourceRegex.MatchString(r) {
				continue
			}
			matches := resourceRegex.FindAllStringSubmatch(r, -1)
			p.Resources = append(p.Resources, Resource{
				Url: matches[0][3],
				Alt: matches[0][2],
			})
		}
	}
	furtherReading := strings.TrimSpace(string(projectFileBytes[furtherReadingStart:]))
	for _, fr := range strings.Split(furtherReading, "\n") {
		if !resourceRegex.MatchString(fr) {
			continue
		}
		matches := resourceRegex.FindAllStringSubmatch(fr, -1)
		p.FurtherReading = append(p.FurtherReading, Resource{
			Url: matches[0][3],
			Alt: matches[0][2],
		})
	}

	return &p, nil
}
