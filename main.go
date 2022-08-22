package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mnys176/md-secretary/compress"
	"github.com/mnys176/md-secretary/contents"
	"github.com/mnys176/md-secretary/create"
	"github.com/mnys176/md-secretary/extend"
	"github.com/mnys176/md-secretary/ingest"
	"github.com/mnys176/md-secretary/scrap"
)

func main() {
	input := os.Args[1:]

	// check if no command is specified
	if len(input) == 0 {
		fmt.Println("No command specified.")
		fmt.Println(globalHelp())
		return
	}

	switch cmd := input[0]; cmd {
	case "contents":
		executable, err := contents.Parse(input)
		if err != nil {
			fmt.Println(err)
			fmt.Println(contents.Help())
			return
		}
		contents.Exec(&executable)
	case "create":
		executable, err := create.Parse(input)
		if err != nil {
			fmt.Println(err)
			fmt.Println(create.Help())
			return
		}
		create.Exec(&executable)
	case "extend":
		executable, err := extend.Parse(input)
		if err != nil {
			fmt.Println(err)
			fmt.Println(extend.Help())
			return
		}
		extend.Exec(&executable)
	case "scrap":
		executable, err := scrap.Parse(input)
		if err != nil {
			fmt.Println(err)
			fmt.Println(scrap.Help())
			return
		}
		scrap.Exec(&executable)
	case "ingest":
		executable, err := ingest.Parse(input)
		if err != nil {
			fmt.Println(err)
			fmt.Println(ingest.Help())
			return
		}
		ingest.Exec(&executable)
	case "compress":
		executable, err := compress.Parse(input)
		if err != nil {
			fmt.Println(err)
			fmt.Println(compress.Help())
			return
		}
		compress.Exec(&executable)
	case "help":
		fmt.Println(globalHelp())
	default:
		fmt.Printf("Invalid command: `%s`\n", cmd)
		fmt.Println(globalHelp())
	}
}

func globalHelp() string {
	wd, _ := os.Getwd()
	data, err := os.ReadFile(filepath.Join(wd, "usage.txt"))
	if err != nil {
		panic(err)
	}
	return string(data)
}
