package extend

import (
	"fmt"
	"os"
	"path/filepath"
)

const defaultNotebookRoot string = "."

type Extend struct {
	ProjectName string
	Path        string
	Help        bool
}

func Build(input []string) (Extend, error) {
	// handle `md-secretary <command>` or `md-secretary <command> --help`
	if len(input) == 1 || len(input) == 2 && (input[1] == "-h" || input[1] == "--help") {
		return Extend{Help: true}, nil
	}

	// configuration variables with defaults
	parsedExtend := Extend{
		ProjectName: input[len(input)-1],
		Path:        defaultNotebookRoot,
	}

	// check if default behavior is desired (no options)
	if len(input) == 2 {
		return parsedExtend, nil
	}

	var addNext bool
	var previous string
	found := map[string]bool{"path": false, "help": false}
	for _, token := range input[1 : len(input)-1] {
		// add values to key-value pair options
		if addNext {
			switch previous {
			case "path":
				parsedExtend.Path = token
			}
			addNext = false
			continue
		}

		switch token {
		case "-h", "--help":
			if !found["help"] {
				found["help"] = true
				parsedExtend.Help = true
			}
		case "-p", "--path":
			if !found["path"] {
				found["path"] = true
				previous = "path"
				addNext = true
			}
		default:
			return parsedExtend, fmt.Errorf("Unknown option: `%s`", token)
		}
	}
	return parsedExtend, nil
}

func Exec(c *Extend) {
	if c.Help {
		Help()
	}
}

func Help() string {
	wd, _ := os.Getwd()
	data, err := os.ReadFile(filepath.Join(wd, "extend", "usage.txt"))
	if err != nil {
		panic(err)
	}
	return string(data)
}
