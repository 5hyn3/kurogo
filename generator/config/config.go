package config

import "io"

type Config struct {
	GlobalParameters map[string]string
	Tasks            []Task
}

type Task struct {
	TaskName     string
	TemplateFrom string
	DeployTo     string
}

type Parser interface {
	Parse(in io.Reader) (Config, error)
}
