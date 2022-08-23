package compress

import "fmt"

func Handle(e *Compress) {
	if e.Help {
		fmt.Println(Help())
		return
	}
	fmt.Println(e)
}
