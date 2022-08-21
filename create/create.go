package create

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const defaultNotebookRoot string = "."

type Create struct {
	ProjectName string
	Path        string
	Help        bool
}

func Build(input []string) (Create, error) {
	// handle `md-secretary <command>` or `md-secretary <command> --help`
	if len(input) == 1 || len(input) == 2 && (input[1] == "-h" || input[1] == "--help") {
		return Create{Help: true}, nil
	}

	// only `--help` option is valid without arguments
	if len(input) == 2 && input[1] != "-h" && input[1] != "--help" && strings.HasPrefix(input[1], "-") {
		return Create{}, fmt.Errorf("Unknown option: `%s`", input[1])
	}

	// configuration variables with defaults
	parsedCreate := Create{
		ProjectName: input[len(input)-1],
		Path:        defaultNotebookRoot,
	}

	// check if default behavior is desired (no options)
	if len(input) == 2 {
		return parsedCreate, nil
	}

	var addNext bool
	var previous string
	found := map[string]bool{"path": false, "help": false}
	for _, token := range input[1 : len(input)-1] {
		// add values to key-value pair options
		if addNext {
			switch previous {
			case "path":
				parsedCreate.Path = token
			}
			addNext = false
			continue
		}

		switch token {
		case "-h", "--help":
			if !found["help"] {
				found["help"] = true
				parsedCreate.Help = true
			}
		case "-p", "--path":
			if !found["path"] {
				found["path"] = true
				previous = "path"
				addNext = true
			}
		default:
			return Create{}, fmt.Errorf("Unknown option: `%s`", token)
		}
	}
	return parsedCreate, nil
}

func Exec(c *Create) {
	if c.Help {
		Help()
	}
}

func Help() string {
	wd, _ := os.Getwd()
	data, err := os.ReadFile(filepath.Join(wd, "create", "usage.txt"))
	if err != nil {
		panic(err)
	}
	return string(data)
}
