package scrap

import "fmt"

func Handle(e *Scrap) {
	if e.Help {
		fmt.Println(Usage)
		return
	}
	fmt.Println(e)
}
