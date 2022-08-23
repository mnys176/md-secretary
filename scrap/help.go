package scrap

import (
	"os"
	"path/filepath"
)

func Help() string {
	usagePath := filepath.Join("scrap", "usage.txt")
	data, err := os.ReadFile(usagePath)
	if err != nil {
		panic(err)
	}
	return string(data)
}
