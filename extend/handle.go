package extend

import "fmt"

func Handle(e *Extend) {
	if e.Help {
		fmt.Println(Usage)
		return
	}
	fmt.Println(e)
}
