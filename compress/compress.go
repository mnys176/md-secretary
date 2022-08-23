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
	Output      string
	Transfer    bool
	Help        bool
}

func (e *Compress) String() string {
	return fmt.Sprintf(
		"Project Name: %s\nPath        : %s\nOutput      : %s\nTransfer    : %t\nHelp        : %t",
		e.ProjectName,
		e.Path,
		e.Output,
		e.Transfer,
		e.Help,
	)
}
