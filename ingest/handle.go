package ingest

import "fmt"

func Handle(e *Ingest) {
	if e.Help {
		fmt.Println(Usage)
		return
	}
	fmt.Println(e)
}
