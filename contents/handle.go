package contents

import "fmt"

func Handle(e *Contents) {
	if e.Help {
		fmt.Println(Help())
		return
	}
	fmt.Println(e)
}
