package create

import (
	"fmt"
	"time"

	"github.com/mnys176/md-secretary/project"
	"github.com/mnys176/md-secretary/utils"
)

func Handle(e *Create) {
	if e.Help {
		fmt.Println(Usage)
		return
	}
	fmt.Println(e)

	year, month, day := time.Now().Date()
	p := project.Project{
		Title:       e.ProjectTitle,
		SystemTitle: utils.Systemify(e.ProjectTitle),
		Abstract:    "",
		Start:       time.Date(year, month, day, 0, 0, 0, 0, time.UTC),
		End:         time.Date(year, month, day, 0, 0, 0, 0, time.UTC),
	}

	err := p.Build(e.Path)
	if err != nil {
		panic(err)
	}
}
