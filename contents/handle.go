package contents

import "fmt"

func Handle(e *Contents) {
	if e.Help {
		fmt.Println(Usage)
		return
	}
	fmt.Println(e)
}
