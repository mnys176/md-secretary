package compress

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/mnys176/md-secretary/config"
)

func Parse(input []string) (Compress, error) {
	// handle `md-secretary <command>` or `md-secretary <command> --help`
	if len(input) == 1 || len(input) == 2 && (input[1] == "-h" || input[1] == "--help") {
		return Compress{Help: true}, nil
	}

	// only `--help` option is valid without arguments
	if len(input) == 2 && input[1] != "-h" && input[1] != "--help" && strings.HasPrefix(input[1], "-") {
		return Compress{}, fmt.Errorf("Unknown option: `%s`", input[1])
	}

	// configuration variables with defaults
	cfg := config.Defaults()
	absNotebookPath, _ := filepath.Abs(cfg.Notebook.Path)
	absCompressionPath, _ := filepath.Abs(cfg.Compression.Path)
	parsedCompress := Compress{
		ProjectTitle: input[len(input)-1],
		Path:         absNotebookPath,
		Output:       absCompressionPath,
	}

	// check if default behavior is desired (no options)
	if len(input) == 2 {
		return parsedCompress, nil
	}

	var addNext bool
	var previous string
	found := map[string]bool{"path": false, "output": false, "transfer": false, "help": false}
	for _, token := range input[1 : len(input)-1] {
		// add values to key-value pair options
		if addNext {
			switch previous {
			case "path":
				absPath, _ := filepath.Abs(token)
				parsedCompress.Path = absPath
			case "output":
				absPath, _ := filepath.Abs(token)
				parsedCompress.Output = absPath
			case "config":
				absPath, _ := filepath.Abs(token)
				parsedCompress.Config = absPath
			}
			addNext = false
			continue
		}

		switch token {
		case "-t", "--transfer":
			if !found["transfer"] {
				found["transfer"] = true
				parsedCompress.Transfer = true
			}
		case "-h", "--help":
			if !found["help"] {
				found["help"] = true
				parsedCompress.Help = true
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
		case "-o", "--output":
			if !found["output"] {
				found["output"] = true
				previous = "output"
				addNext = true
			}
		default:
			return Compress{}, fmt.Errorf("Unknown option: `%s`", token)
		}
	}
	return parsedCompress, nil
}
