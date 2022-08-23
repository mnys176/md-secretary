package create

import "fmt"

func Handle(e *Create) {
	if e.Help {
		fmt.Println(Help())
		return
	}
	fmt.Println(e)
}
