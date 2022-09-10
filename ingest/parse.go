package ingest

import (
	"fmt"
	"path/filepath"
	"strings"
)

func Parse(input []string) (*Ingest, error) {
	// handle `md-secretary <command>` or `md-secretary <command> --help`
	if len(input) == 1 || len(input) == 2 && (input[1] == "-h" || input[1] == "--help") {
		return &Ingest{Help: true}, nil
	}

	// only `--help` option is valid without arguments
	if len(input) == 2 && input[1] != "-h" && input[1] != "--help" && strings.HasPrefix(input[1], "-") {
		return nil, fmt.Errorf("Unknown option: `%s`", input[1])
	}

	// stored `ingest` command
	parsedIngest := Ingest{PathToJson: input[len(input)-1]}

	// check if default behavior is desired (no options)
	if len(input) == 2 {
		return &parsedIngest, nil
	}

	var addNext bool
	var previous string
	found := map[string]bool{"path": false, "config": false, "force": false, "help": false}
	for _, token := range input[1 : len(input)-1] {
		// add values to key-value pair options
		if addNext {
			switch previous {
			case "path":
				absPath, _ := filepath.Abs(token)
				parsedIngest.Path = absPath
			case "config":
				absPath, _ := filepath.Abs(token)
				parsedIngest.Config = absPath
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
		case "-c", "--config":
			if !found["config"] {
				found["config"] = true
				previous = "config"
				addNext = true
			}
		default:
			return nil, fmt.Errorf("Unknown option: `%s`", token)
		}
	}
	return &parsedIngest, nil
}
