package ingest

import (
	_ "embed"
	"fmt"
)

//go:embed usage.txt
var Usage string

type Ingest struct {
	PathToJson string
	Path       string
	Config     string
	Force      bool
	Help       bool
}

func (e *Ingest) String() string {
	const template string = `Path to JSON: %s
Path        : %s
Config      : %s
Force       : %t
Help        : %t`

	return fmt.Sprintf(
		template,
		e.PathToJson,
		e.Path,
		e.Config,
		e.Force,
		e.Help,
	)
}
