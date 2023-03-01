package model

type Property interface{}

type Schema struct {
	Properties map[string]Property `yaml:"properties"`
}
