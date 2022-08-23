package scrap

import (
	_ "embed"
	"fmt"
)

//go:embed usage.txt
var Usage string

type Scrap struct {
	ProjectName string
	Path        string
	Force       bool
	Help        bool
}

func (e *Scrap) String() string {
	return fmt.Sprintf(
		"Project Name: %s\nPath        : %s\nForce       : %t\nHelp        : %t",
		e.ProjectName,
		e.Path,
		e.Force,
		e.Help,
	)
}
