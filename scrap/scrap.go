package scrap

import (
	_ "embed"
	"fmt"
)

//go:embed usage.txt
var Usage string

type Scrap struct {
	ProjectTitle string
	Path         string
	Config       string
	Force        bool
	Help         bool
}

func (e *Scrap) String() string {
	const template string = `Project Title: %s
Path         : %s
Config       : %s
Force        : %t
Help         : %t`
	return fmt.Sprintf(
		template,
		e.ProjectTitle,
		e.Path,
		e.Config,
		e.Force,
		e.Help,
	)
}
