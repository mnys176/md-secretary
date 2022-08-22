package ingest

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const defaultNotebookPath string = "."

type Ingest struct {
	PathToJson string
	Path       string
	Force      bool
	Help       bool
}

func (e *Ingest) String() string {
	return fmt.Sprintf(
		"Path to JSON: %s\nPath        : %s\nForce       : %t\nHelp        : %t",
		e.PathToJson,
		e.Path,
		e.Force,
		e.Help,
	)
}

func Parse(input []string) (Ingest, error) {
	// handle `md-secretary <command>` or `md-secretary <command> --help`
	if len(input) == 1 || len(input) == 2 && (input[1] == "-h" || input[1] == "--help") {
		return Ingest{Help: true}, nil
	}

	// only `--help` option is valid without arguments
	if len(input) == 2 && input[1] != "-h" && input[1] != "--help" && strings.HasPrefix(input[1], "-") {
		return Ingest{}, fmt.Errorf("Unknown option: `%s`", input[1])
	}

	// configuration variables with defaults
	absNotebookPath, _ := filepath.Abs(defaultNotebookPath)
	parsedIngest := Ingest{
		PathToJson: input[len(input)-1],
		Path:       absNotebookPath,
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
				absPath, _ := filepath.Abs(token)
				parsedIngest.Path = absPath
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

func Exec(e *Ingest) {
	if e.Help {
		fmt.Println(Help())
		return
	}
	fmt.Println(e)
}

func Help() string {
	usagePath := filepath.Join("ingest", "usage.txt")
	data, err := os.ReadFile(usagePath)
	if err != nil {
		panic(err)
	}
	return string(data)
}
