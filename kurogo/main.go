package main

import (
	"errors"
	"github.com/5hyn3/kurogo/generator"
	"github.com/5hyn3/kurogo/generator/config/parser"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "kurogo"
	app.Usage = "A simple template processor by golang"
	app.Version = "0.0.1"
	app.EnableBashCompletion = true
	app.Action = func(c *cli.Context) error {
		if c.NArg() != 1 {
			return errors.New("the number of parameters is not one")
		}
		file, err := os.Open(c.Args()[0])
		if err != nil {
			panic(err)
		}
		defer file.Close()

		generator.Generator(file, parser.NewYml())
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
