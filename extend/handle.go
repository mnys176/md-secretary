package extend

import "fmt"

func Handle(e *Extend) {
	if e.Help {
		fmt.Println(Help())
		return
	}
	fmt.Println(e)
}
