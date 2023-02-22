package model

type Model struct {
	Name   string                 `yaml:"name"`
	Schema map[string]interface{} `yaml:"schema"`
}
