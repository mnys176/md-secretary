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
		showGlobalHelp()
		return
	}

	switch cmd := input[0]; cmd {
	case "contents":
		executable, err := contents.Build(input)
		if err != nil {
			fmt.Println(err)
			fmt.Println(contents.Help())
		}
		contents.Exec(&executable)
	case "create":
		executable, err := create.Build(input)
		if err != nil {
			fmt.Println(err)
			fmt.Println(create.Help())
		}
		create.Exec(&executable)
	case "extend":
		executable, err := extend.Build(input)
		if err != nil {
			fmt.Println(err)
			fmt.Println(extend.Help())
		}
		extend.Exec(&executable)
	case "scrap":
		executable, err := scrap.Build(input)
		if err != nil {
			fmt.Println(err)
			fmt.Println(scrap.Help())
		}
		scrap.Exec(&executable)
	case "ingest":
		executable, err := ingest.Build(input)
		if err != nil {
			fmt.Println(err)
			fmt.Println(ingest.Help())
		}
		ingest.Exec(&executable)
	case "compress":
		executable, err := compress.Build(input)
		if err != nil {
			fmt.Println(err)
			fmt.Println(compress.Help())
		}
		compress.Exec(&executable)
	case "help":
		showGlobalHelp()
	default:
		fmt.Printf("Invalid command: `%s`\n", cmd)
		showGlobalHelp()
	}
}

func showGlobalHelp() {
	wd, _ := os.Getwd()
	data, err := os.ReadFile(filepath.Join(wd, "usage.txt"))
	if err != nil {
		panic(err)
	}
	os.Stdout.Write(data)
}
