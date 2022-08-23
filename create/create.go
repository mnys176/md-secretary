package create

import (
	_ "embed"
	"fmt"
)

//go:embed usage.txt
var Usage string

type Create struct {
	ProjectName string
	Path        string
	Help        bool
}

func (e *Create) String() string {
	return fmt.Sprintf(
		"Project Name: %s\nPath        : %s\nHelp        : %t",
		e.ProjectName,
		e.Path,
		e.Help,
	)
}
