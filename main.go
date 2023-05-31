package main

import (
	"fmt"
	"os"

	"github.com/Pedrommb91/go-img-scrapper/internal/args"
	"github.com/Pedrommb91/go-img-scrapper/internal/commands"
	"github.com/Pedrommb91/go-img-scrapper/internal/commands/scrapers"
)

func main() {
	parser := args.NewArgsParser(os.Args)
	err := parser.Validate()
	if err != nil {
		fmt.Println(err)
		commands.DisplayHelp()
		os.Exit(1)
	}

	args, err := parser.Parse()
	if err != nil {
		fmt.Println(err)
		commands.DisplayHelp()
		os.Exit(1)
	}

	if args.Help {
		commands.DisplayHelp()
		os.Exit(0)
	}

	for _, s := range args.Scrape {
		err := scrapers.NewHTMLScraper(s).Scrape()
		if err != nil {
			fmt.Println(err)
		}
	}
}
