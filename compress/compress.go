package compress

import (
	_ "embed"
	"fmt"
)

//go:embed usage.txt
var Usage string

type Compress struct {
	ProjectTitle string
	Path         string
	Config       string
	Output       string
	Transfer     bool
	Force        bool
	Help         bool
}

func (e *Compress) String() string {
	const template string = `Project Title: %s
Path         : %s
Output       : %s
Config       : %s
Transfer     : %t
Force        : %t
Help         : %t`

	return fmt.Sprintf(
		template,
		e.ProjectTitle,
		e.Path,
		e.Output,
		e.Config,
		e.Transfer,
		e.Help,
	)
}
