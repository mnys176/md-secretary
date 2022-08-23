package scrap

import "fmt"

func Handle(e *Scrap) {
	if e.Help {
		fmt.Println(Help())
		return
	}
	fmt.Println(e)
}
