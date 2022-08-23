package compress

import "fmt"

func Handle(e *Compress) {
	if e.Help {
		fmt.Println(Usage)
		return
	}
	fmt.Println(e)
}
