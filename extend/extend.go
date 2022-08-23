package extend

import "fmt"

type Extend struct {
	ProjectName string
	Path        string
	Help        bool
}

func (e *Extend) String() string {
	return fmt.Sprintf(
		"Project Name: %s\nPath        : %s\nHelp        : %t",
		e.ProjectName,
		e.Path,
		e.Help,
	)
}
