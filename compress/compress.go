package compress

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	defaultNotebookRoot    string = "."
	defaultOutputDirectory string = "."
)

type Compress struct {
	ProjectName string
	Path        string
	Output      string
	Transfer    bool
	Help        bool
}

func Build(input []string) (Compress, error) {
	// handle `md-secretary <command>` or `md-secretary <command> --help`
	if len(input) == 1 || len(input) == 2 && (input[1] == "-h" || input[1] == "--help") {
		return Compress{Help: true}, nil
	}

	// configuration variables with defaults
	parsedCompress := Compress{
		ProjectName: input[len(input)-1],
		Path:        defaultNotebookRoot,
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
				parsedCompress.Path = token
			case "output":
				parsedCompress.Output = token
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
		case "-o", "--output":
			if !found["output"] {
				found["output"] = true
				previous = "output"
				addNext = true
			}
		default:
			return parsedCompress, fmt.Errorf("Unknown option: `%s`\n", token)
		}
	}
	return parsedCompress, nil
}

func Exec(c *Compress) {
	if c.Help {
		Help()
	}
}

func Help() {
	wd, _ := os.Getwd()
	data, err := os.ReadFile(filepath.Join(wd, "compress", "usage.txt"))
	if err != nil {
		panic(err)
	}
	os.Stdout.Write(data)
}
