package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const prompt string = "Proceed? [yN]: "

func Confirm(dialog string) bool {
	stdin := bufio.NewScanner(os.Stdin)
	fmt.Print("\n" + dialog + "\n\n" + prompt)
	stdin.Scan()
	return strings.ToUpper(stdin.Text()) == "Y"
}
