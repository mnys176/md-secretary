package contents

import "fmt"

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
