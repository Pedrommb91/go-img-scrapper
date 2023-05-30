package args

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Pedrommb91/go-img-scrapper/internal/args/validators"
	"github.com/Pedrommb91/go-img-scrapper/pkg/slices"
)

type argsParser struct {
	args   []string
	parsed *Args
}

func NewArgsParser(args []string) argsParser {
	return argsParser{args: args}
}

func (p argsParser) Validate() error {
	if len(p.args) < 2 {
		return errors.New("error: url must be declared")
	}

	return nil
}

func (p argsParser) Parse() (*Args, error) {
	p.parsed = &Args{}
	_, hasHelp := slices.Contains(p.args, func(arg string) bool {
		return arg == "--help" || arg == "-h"
	})
	p.parsed.Help = hasHelp

	i, hasFile := slices.Contains(p.args, func(arg string) bool {
		return arg == "--file" || arg == "-f"
	})
	if hasFile {
		if len(p.args) <= i {
			return nil, fmt.Errorf("error: must set a file path after the file flag")
		}
		if err := p.ParseFile(p.args[i+1]); err != nil {
			return nil, err
		}

	}

	return p.parsed, nil
}

func (p argsParser) ParseFile(file string) error {
	v := validators.NewArgFileValidator(file)
	err := v.Validate()
	if err != nil {
		return err
	}

	jsonFile, err := os.Open(file)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(byteValue, &p.parsed.Scrape); err != nil {
		return err
	}

	p.parsed.File = file
	return nil
}
