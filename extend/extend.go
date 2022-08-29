package extend

import (
	_ "embed"
	"fmt"
)

//go:embed usage.txt
var Usage string

type Extend struct {
	ProjectTitle string
	Path         string
	Config       string
	Help         bool
}

func (e *Extend) String() string {
	const template string = `Project Title: %s
Path         : %s
Config       : %s
Help         : %t`

	return fmt.Sprintf(
		template,
		e.ProjectTitle,
		e.Path,
		e.Config,
		e.Help,
	)
}
