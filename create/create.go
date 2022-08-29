package create

import (
	_ "embed"
	"fmt"
)

//go:embed usage.txt
var Usage string

type Create struct {
	ProjectTitle string
	Path         string
	Config       string
	Help         bool
}

func (e Create) String() string {
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
