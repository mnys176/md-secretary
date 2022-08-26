package compress

import (
	_ "embed"
	"fmt"
)

//go:embed usage.txt
var Usage string

type Compress struct {
	ProjectName string
	Path        string
	Config      string
	Output      string
	Transfer    bool
	Help        bool
}

func (e *Compress) String() string {
	const template string = `Project Name: %s
Path        : %s
Output      : %s
Config      : %s
Transfer    : %t
Help        : %t`

	return fmt.Sprintf(
		template,
		e.ProjectName,
		e.Path,
		e.Output,
		e.Config,
		e.Transfer,
		e.Help,
	)
}
