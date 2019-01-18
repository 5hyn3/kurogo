package generator

import (
	"github.com/5hyn3/kurogo/generator/config"
	"github.com/5hyn3/kurogo/generator/template"
	"io"
	"os"
)

func Generator(in io.Reader, parser config.Parser) {
	conf, err := parser.Parse(in)
	if err != nil {
		panic(err)
	}

	tasks := conf.Tasks

	for _, t := range tasks {
		func() {
			templateFrom, err := os.Open(t.TemplateFrom)
			defer templateFrom.Close()
			if err != nil {
				panic(err)
			}
			deployTo, err := os.OpenFile(t.DeployTo, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
			defer deployTo.Close()
			if err != nil {
				panic(err)
			}
			err = template.Processor(templateFrom, deployTo, conf.GlobalParameters)
			if err != nil {
				panic(err)
			}
		}()
	}
}
