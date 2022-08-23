package ingest

import "fmt"

type Ingest struct {
	PathToJson string
	Path       string
	Force      bool
	Help       bool
}

func (e *Ingest) String() string {
	return fmt.Sprintf(
		"Path to JSON: %s\nPath        : %s\nForce       : %t\nHelp        : %t",
		e.PathToJson,
		e.Path,
		e.Force,
		e.Help,
	)
}
