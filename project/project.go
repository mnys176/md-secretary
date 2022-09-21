package project

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"github.com/mnys176/md-secretary/config"
	"github.com/mnys176/md-secretary/utils"
)

type (
	Project struct {
		Title          string
		SystemTitle    string
		Abstract       string
		Resources      []Resource
		FurtherReading []Resource
		Start          *Marker
		End            *Marker
		Markers        []Marker
		Notebook       string
	}
	ProjectTemplateData struct {
		Title          string
		Abstract       string
		Resources      string
		FurtherReading string
	}
	File struct {
		FileName string `json:"fileName"`
		Data     []byte `json:"data"`
	}
	Resource struct {
		Url string `json:"url"`
		Alt string `json:"alt"`
	}
)

/*

type Project struct {
	Title string
	SystemTitle string
	Abstract string
	Media []struct {
		Filename string
		Base64 string
	}
	Resources []Note
	FurtherReading []Note
	Markers []Marker
}

type Marker struct {
	Date string
	LogFile []BrainDump
	SummaryFile []Milestone
}

type BrainDump struct {
	Date string
	Content []Note
}

type Milestone struct {
	Summary string
	BrainDump
}

type Note struct {
	Text string
	Urls []struct {
		Link string
		Alt string
	}
	Children []Note
}

*/

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

func (p Project) Build(cfg *config.Config) error {
	projectPath := filepath.Join(p.Notebook, p.SystemTitle)
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
		Abstract:       cfg.Project.AbstractTemplate,
		Resources:      cfg.Project.ResourcesTemplate,
		FurtherReading: cfg.Project.FurtherReadingTemplate,
	})
	if err != nil {
		return err
	}

	p.Start.Compact = cfg.Notebook.CompactMarkerDirectory == "compact"

	// build first marker
	err = p.Start.Build(cfg)
	if err != nil {
		return err
	}
	return nil
}

func (p Project) Append(cfg *config.Config) error {
	projectPath := filepath.Join(p.Notebook, p.SystemTitle)
	m := NewMarker(projectPath)
	m.Compact = cfg.Notebook.CompactMarkerDirectory == "compact"
	err := m.Build(cfg)
	if err != nil {
		return err
	}
	return nil
}

func (p Project) TearDown(force bool, cfg *config.Config) error {
	projectPath := filepath.Join(p.Notebook, p.SystemTitle)
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

func (p *Project) Export(outputPath string, transfer bool, cfg *config.Config) error {
	var jsonTitle string
	if filepath.Ext(outputPath) == ".json" {
		jsonTitle = filepath.Base(outputPath)
		outputPath = filepath.Dir(outputPath)
	} else {
		jsonTitle = strings.TrimSpace(cfg.Compression.JsonTitle)
	}
	jsonTitle = strings.TrimSuffix(jsonTitle, ".json")
	jsonTitle = strings.TrimSpace(jsonTitle)
	jsonTitle = utils.Systemify(jsonTitle)

	// replace `$project` with the sanitized project title
	jsonTitle = strings.ReplaceAll(jsonTitle, "$project", p.SystemTitle)

	// replace `$date` with the current date
	year, month, day := time.Now().Date()
	today := fmt.Sprintf("%02d-%02d-%04d", month, day, year)
	jsonTitle = strings.ReplaceAll(jsonTitle, "$date", today)

	// NOTE: Anything created by a command is "systemified". Since
	//       the parent directory is assumed to exist prior to the
	//       creation of the JSON file, no sanitation will be
	//       performed on the parent directory, but sanitation will
	//       always be done on the filename itself.
	jsonFilePath := filepath.Join(outputPath, jsonTitle+".json")

	// create JSON file
	jsonFile, err := os.Create(jsonFilePath)
	defer jsonFile.Close()
	if err != nil {
		return err
	}

	b, err := json.MarshalIndent(p, "", "    ")
	if err != nil {
		return err
	}
	os.Stdout.Write(b)
	return nil
}

func (p Project) MarshalJSON() ([]byte, error) {
	type (
		ProjectJson struct {
			Title          string     `json:"title"`
			SystemTitle    string     `json:"systemTitle"`
			Abstract       string     `json:"abstract"`
			Start          int64      `json:"start"`
			End            int64      `json:"end"`
			Resources      []Resource `json:"resources"`
			FurtherReading []Resource `json:"furtherReading"`
			ProjectFile    File       `json:"projectFile"`
			Markers        []Marker   `json:"markers"`
			Media          []File     `json:"media"`
		}
	)

	projectPath := filepath.Join(p.Notebook, p.SystemTitle)
	projectFilePath := filepath.Join(projectPath, p.SystemTitle+".md")
	mediaPath := filepath.Join(projectPath, "media")

	// compress and encode project file
	projectFileBytes, err := os.ReadFile(projectFilePath)
	if err != nil {
		return nil, err
	}
	projectFileCompressed, err := utils.CompressEncode(projectFileBytes)
	if err != nil {
		return nil, err
	}

	// compress and encode media
	media := []File{}
	mediaFiles, err := os.ReadDir(mediaPath)
	if err != nil {
		return nil, err
	}
	for _, f := range mediaFiles {
		fBytes, err := os.ReadFile(filepath.Join(mediaPath, f.Name()))
		if err != nil {
			return nil, err
		}
		fCompressed, err := utils.CompressEncode(fBytes)
		if err != nil {
			return nil, err
		}
		media = append(media, File{
			FileName: f.Name(),
			Data: fCompressed,
		})
	}

	return json.Marshal(ProjectJson{
		Title:          p.Title,
		SystemTitle:    p.SystemTitle,
		Abstract:       p.Abstract,
		Start:          p.Start.Date.Unix(),
		End:            p.End.Date.Unix(),
		Resources:      p.Resources,
		FurtherReading: p.FurtherReading,
		ProjectFile:    File{p.SystemTitle + ".md", projectFileCompressed},
		Markers:        p.Markers,
		Media:          media,
	})
}
