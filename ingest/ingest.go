package ingest

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const defaultNotebookRoot string = "."

type Ingest struct {
	PathToJson string
	Path       string
	Force      bool
	Help       bool
}

func Build(input []string) (Ingest, error) {
	// handle `md-secretary <command>` or `md-secretary <command> --help`
	if len(input) == 1 || len(input) == 2 && (input[1] == "-h" || input[1] == "--help") {
		return Ingest{Help: true}, nil
	}

	// only `--help` option is valid without arguments 
	if len(input) == 2 && input[1] != "-h" && input[1] != "--help" && strings.HasPrefix(input[1], "-") {
		return Ingest{}, fmt.Errorf("Unknown option: `%s`", input[1])
	}

	// configuration variables with defaults
	parsedIngest := Ingest{
		PathToJson: input[len(input)-1],
		Path:       defaultNotebookRoot,
	}

	// check if default behavior is desired (no options)
	if len(input) == 2 {
		return parsedIngest, nil
	}

	var addNext bool
	var previous string
	found := map[string]bool{"path": false, "force": false, "help": false}
	for _, token := range input[1 : len(input)-1] {
		// add values to key-value pair options
		if addNext {
			switch previous {
			case "path":
				parsedIngest.Path = token
			}
			addNext = false
			continue
		}

		switch token {
		case "-f", "--force":
			if !found["force"] {
				found["force"] = true
				parsedIngest.Force = true
			}
		case "-h", "--help":
			if !found["help"] {
				found["help"] = true
				parsedIngest.Help = true
			}
		case "-p", "--path":
			if !found["path"] {
				found["path"] = true
				previous = "path"
				addNext = true
			}
		default:
			return Ingest{}, fmt.Errorf("Unknown option: `%s`", token)
		}
	}
	return parsedIngest, nil
}

func Exec(c *Ingest) {
	if c.Help {
		Help()
	}
}

func Help() string {
	wd, _ := os.Getwd()
	data, err := os.ReadFile(filepath.Join(wd, "ingest", "usage.txt"))
	if err != nil {
		panic(err)
	}
	return string(data)
}
