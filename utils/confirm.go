package utils

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

const prompt string = "Proceed? [yN]: "

func Confirm(dialog string) bool {
	stdin := bufio.NewScanner(os.Stdin)
	fmt.Print(dialog + "\n" + prompt)
	stdin.Scan()
	return strings.ToUpper(stdin.Text()) == "Y"
}
