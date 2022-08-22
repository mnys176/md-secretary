package contents

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const defaultNotebookPath string = "."

type Contents struct {
	ProjectName string
	Path        string
	Help        bool
}

func (e *Contents) String() string {
	return fmt.Sprintf(
		"Project Name: %s\nPath        : %s\nHelp        : %t",
		e.ProjectName,
		e.Path,
		e.Help,
	)
}

func Parse(input []string) (Contents, error) {
	// handle `md-secretary <command>` or `md-secretary <command> --help`
	if len(input) == 1 || len(input) == 2 && (input[1] == "-h" || input[1] == "--help") {
		return Contents{Help: true}, nil
	}

	// only `--help` option is valid without arguments
	if len(input) == 2 && input[1] != "-h" && input[1] != "--help" && strings.HasPrefix(input[1], "-") {
		return Contents{}, fmt.Errorf("Unknown option: `%s`", input[1])
	}

	// configuration variables with defaults
	absNotebookPath, _ := filepath.Abs(defaultNotebookPath)
	parsedContents := Contents{
		ProjectName: input[len(input)-1],
		Path:        absNotebookPath,
	}

	// check if default behavior is desired (no options)
	if len(input) == 2 {
		return parsedContents, nil
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
				parsedContents.Path = absPath
			}
			addNext = false
			continue
		}

		switch token {
		case "-h", "--help":
			if !found["help"] {
				found["help"] = true
				parsedContents.Help = true
			}
		case "-p", "--path":
			if !found["path"] {
				found["path"] = true
				previous = "path"
				addNext = true
			}
		default:
			return Contents{}, fmt.Errorf("Unknown option: `%s`", token)
		}
	}
	return parsedContents, nil
}

func Exec(e *Contents) {
	if e.Help {
		fmt.Println(Help())
		return
	}
	fmt.Println(e)
}

func Help() string {
	usagePath := filepath.Join("contents", "usage.txt")
	data, err := os.ReadFile(usagePath)
	if err != nil {
		panic(err)
	}
	return string(data)
}
