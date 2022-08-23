package create

import "fmt"

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
