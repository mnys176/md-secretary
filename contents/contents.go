package contents

import (
	_ "embed"
	"fmt"
)

//go:embed usage.txt
var Usage string

type Contents struct {
	ProjectTitle string
	Path         string
	Config       string
	Help         bool
}

func (e *Contents) String() string {
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
