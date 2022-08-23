package ingest

import "fmt"

func Handle(e *Ingest) {
	if e.Help {
		fmt.Println(Help())
		return
	}
	fmt.Println(e)
}
