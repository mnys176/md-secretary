package main

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/mnys176/md-secretary/compress"
	"github.com/mnys176/md-secretary/contents"
	"github.com/mnys176/md-secretary/create"
	"github.com/mnys176/md-secretary/extend"
	"github.com/mnys176/md-secretary/ingest"
	"github.com/mnys176/md-secretary/scrap"
)

//go:embed usage.txt
var Usage string

func main() {
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
		err = contents.Handle(executable)
		if err != nil {
			fmt.Println(err)
		}
	case "create":
		executable, err := create.Parse(input)
		if err != nil {
			fmt.Println(err)
			fmt.Println(create.Usage)
			return
		}
		err = create.Handle(executable)
		if err != nil {
			fmt.Println(err)
		}
	case "extend":
		executable, err := extend.Parse(input)
		if err != nil {
			fmt.Println(err)
			fmt.Println(extend.Usage)
			return
		}
		err = extend.Handle(executable)
		if err != nil {
			fmt.Println(err)
		}
	case "scrap":
		executable, err := scrap.Parse(input)
		if err != nil {
			fmt.Println(err)
			fmt.Println(scrap.Usage)
			return
		}
		err = scrap.Handle(executable)
		if err != nil {
			fmt.Println(err)
		}
	case "ingest":
		executable, err := ingest.Parse(input)
		if err != nil {
			fmt.Println(err)
			fmt.Println(ingest.Usage)
			return
		}
		err = ingest.Handle(executable)
		if err != nil {
			fmt.Println(err)
		}
	case "compress":
		executable, err := compress.Parse(input)
		if err != nil {
			fmt.Println(err)
			fmt.Println(compress.Usage)
			return
		}
		err = compress.Handle(executable)
		if err != nil {
			fmt.Println(err)
		}
	case "help":
		fmt.Println(Usage)
	default:
		fmt.Printf("Invalid command: `%s`\n", cmd)
		fmt.Println(Usage)
	}
}
