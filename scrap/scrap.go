package scrap

import (
	"fmt"
	"os"
	"path/filepath"
)

const defaultNotebookRoot string = "."

type Scrap struct {
	ProjectName string
	Path        string
	Force       bool
	Help        bool
}

func Build(input []string) (Scrap, error) {
	// handle `md-secretary <command>` or `md-secretary <command> --help`
	if len(input) == 1 || len(input) == 2 && (input[1] == "-h" || input[1] == "--help") {
		return Scrap{Help: true}, nil
	}

	// configuration variables with defaults
	parsedScrap := Scrap{
		ProjectName: input[len(input)-1],
		Path:        defaultNotebookRoot,
	}

	// check if default behavior is desired (no options)
	if len(input) == 2 {
		return parsedScrap, nil
	}

	var addNext bool
	var previous string
	found := map[string]bool{"path": false, "force": false, "help": false}
	for _, token := range input[1 : len(input)-1] {
		// add values to key-value pair options
		if addNext {
			switch previous {
			case "path":
				parsedScrap.Path = token
			}
			addNext = false
			continue
		}

		switch token {
		case "-f", "--force":
			if !found["force"] {
				found["force"] = true
				parsedScrap.Force = true
			}
		case "-h", "--help":
			if !found["help"] {
				found["help"] = true
				parsedScrap.Help = true
			}
		case "-p", "--path":
			if !found["path"] {
				found["path"] = true
				previous = "path"
				addNext = true
			}
		default:
			return parsedScrap, fmt.Errorf("Unknown option: `%s`\n", token)
		}
	}
	return parsedScrap, nil
}

func Exec(c *Scrap) {
	if c.Help {
		Help()
	}
}

func Help() {
	wd, _ := os.Getwd()
	data, err := os.ReadFile(filepath.Join(wd, "scrap", "usage.txt"))
	if err != nil {
		panic(err)
	}
	os.Stdout.Write(data)
}
