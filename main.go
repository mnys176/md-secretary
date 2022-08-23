package main

import (
	"fmt"
	"os"
	// "path/filepath"
	_ "embed"

	"github.com/mnys176/md-secretary/compress"
	"github.com/mnys176/md-secretary/config"
	"github.com/mnys176/md-secretary/contents"
	"github.com/mnys176/md-secretary/create"
	"github.com/mnys176/md-secretary/extend"
	"github.com/mnys176/md-secretary/ingest"
	"github.com/mnys176/md-secretary/scrap"
)

//go:embed usage.txt
var Usage string

func main() {
	config.Foo()
	input := os.Args[1:]

	// check if no command is specified
	if len(input) == 0 {
		fmt.Println("No command specified.")
		fmt.Println(Usage)
		return
	}

	switch cmd := input[0]; cmd {
	case "contents":
		executable, err := contents.Parse(input)
		if err != nil {
			fmt.Println(err)
			fmt.Println(contents.Usage)
			return
		}
		contents.Handle(&executable)
	case "create":
		executable, err := create.Parse(input)
		if err != nil {
			fmt.Println(err)
			fmt.Println(create.Usage)
			return
		}
		create.Handle(&executable)
	case "extend":
		executable, err := extend.Parse(input)
		if err != nil {
			fmt.Println(err)
			fmt.Println(extend.Usage)
			return
		}
		extend.Handle(&executable)
	case "scrap":
		executable, err := scrap.Parse(input)
		if err != nil {
			fmt.Println(err)
			fmt.Println(scrap.Usage)
			return
		}
		scrap.Handle(&executable)
	case "ingest":
		executable, err := ingest.Parse(input)
		if err != nil {
			fmt.Println(err)
			fmt.Println(ingest.Usage)
			return
		}
		ingest.Handle(&executable)
	case "compress":
		executable, err := compress.Parse(input)
		if err != nil {
			fmt.Println(err)
			fmt.Println(compress.Usage)
			return
		}
		compress.Handle(&executable)
	case "help":
		fmt.Println(Usage)
	default:
		fmt.Printf("Invalid command: `%s`\n", cmd)
		fmt.Println(Usage)
	}
}
