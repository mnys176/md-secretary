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
	Config        string
	Help        bool
}

func (e *Create) String() string {
	const template string = `Project Name: %s
Path        : %s
Config      : %s
Help        : %t`

	return fmt.Sprintf(
		template,
		e.ProjectName,
		e.Path,
		e.Config,
		e.Help,
	)
}
