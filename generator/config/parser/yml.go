package parser

import (
	"errors"
	"github.com/5hyn3/kurogo/generator/config"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
)

var (
	ErrEmptyTasks = errors.New("empty tasks")
)

type YmlConfig struct {
	YmlGlobalParameters []YmlGlobalParameter `yaml:"global_parameters"`
	YmlTasks            []YmlTask `yaml:"tasks"`
}

type YmlGlobalParameter struct {
	Key string `yaml:"key"`
	Value string `yaml:"value"`
}

type YmlTask struct {
	YmlTaskName     string `yaml:"task_name"`
	YmlTemplateFrom string `yaml:"template_from"`
	YmlDeployTo     string `yaml:"deploy_to"`
}

type yamlParser struct {
}

func (p yamlParser) Parse(in io.Reader) (config.Config, error) {
	var conf config.Config
	buf, err := ioutil.ReadAll(in)
	if err != nil {
		return conf, err
	}

	var ymlConfig YmlConfig

	err = yaml.Unmarshal(buf, &ymlConfig)
	if err != nil {
		return conf, err
	}

	if len(ymlConfig.YmlTasks) == 0 {
		return conf, ErrEmptyTasks
	}

	conf, err = ymlConfig.Config()

	if err != nil {
		return conf, err
	}

	return conf, nil
}

func NewYml() config.Parser {
	return yamlParser{}
}

func (c *YmlConfig) Config() (config.Config, error) {
	var globalParameters = make(map[string]string, len(c.YmlGlobalParameters))
	for i := range c.YmlGlobalParameters {
		globalParameters[c.YmlGlobalParameters[i].Key] = c.YmlGlobalParameters[i].Value
	}

	var tasks = make([]config.Task, len(c.YmlTasks), len(c.YmlTasks))
	for i := range c.YmlTasks {
		tasks[i] = config.Task{
			TaskName:     c.YmlTasks[i].YmlTaskName,
			TemplateFrom: c.YmlTasks[i].YmlTemplateFrom,
			DeployTo:     c.YmlTasks[i].YmlDeployTo,
		}
	}
	return config.Config{
		GlobalParameters: globalParameters,
		Tasks:            tasks,
	}, nil
}
