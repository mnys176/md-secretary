package create

import "fmt"

func Handle(e *Create) {
	if e.Help {
		fmt.Println(Usage)
		return
	}
	fmt.Println(e)
}
