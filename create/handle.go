package create

import (
	"fmt"
	"strings"
	"time"

	"github.com/mnys176/md-secretary/project"
)

func Handle(e *Create) {
	if e.Help {
		fmt.Println(Usage)
		return
	}
	fmt.Println(e)

	year, month, _ := time.Now().Date()
	p := project.Project{
		Title:       e.ProjectTitle,
		SystemTitle: strings.ToLower(strings.ReplaceAll(e.ProjectTitle, " ", "-")),
		Abstract:    "",
		Start:       time.Date(year, month, 1, 0, 0, 0, 0, time.UTC),
		End:         time.Date(year, month, 1, 0, 0, 0, 0, time.UTC),
	}

	err := p.Build(e.Path)
	if err != nil {
		panic(err)
	}
}
