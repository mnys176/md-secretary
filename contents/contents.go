package contents

import (
	_ "embed"
	"fmt"
)

//go:embed usage.txt
var Usage string

type Contents struct {
	ProjectName string
	Path        string
	Config      string
	Help        bool
}

func (e *Contents) String() string {
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
