package contents

import (
	"os"
	"path/filepath"
)

func Help() string {
	usagePath := filepath.Join("contents", "usage.txt")
	data, err := os.ReadFile(usagePath)
	if err != nil {
		panic(err)
	}
	return string(data)
}
