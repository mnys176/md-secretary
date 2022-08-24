package extend

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/mnys176/md-secretary/config"
)

func Parse(input []string) (Extend, error) {
	// handle `md-secretary <command>` or `md-secretary <command> --help`
	if len(input) == 1 || len(input) == 2 && (input[1] == "-h" || input[1] == "--help") {
		return Extend{Help: true}, nil
	}

	// only `--help` option is valid without arguments
	if len(input) == 2 && input[1] != "-h" && input[1] != "--help" && strings.HasPrefix(input[1], "-") {
		return Extend{}, fmt.Errorf("Unknown option: `%s`", input[1])
	}

	// configuration variables with defaults
	cfg := config.Defaults()
	absNotebookPath, _ := filepath.Abs(cfg.Notebook.Path)
	parsedExtend := Extend{
		ProjectName: input[len(input)-1],
		Path:        absNotebookPath,
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
				absPath, _ := filepath.Abs(token)
				parsedExtend.Path = absPath
			case "config":
				absPath, _ := filepath.Abs(token)
				parsedExtend.Config = absPath
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
		case "-c", "--config":
			if !found["config"] {
				found["config"] = true
				previous = "config"
				addNext = true
			}
		default:
			return Extend{}, fmt.Errorf("Unknown option: `%s`", token)
		}
	}
	return parsedExtend, nil
}
