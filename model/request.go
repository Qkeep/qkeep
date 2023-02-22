package model

type Request struct {
	Method  string                 `yaml:"method"`
	Path    string                 `yaml:"path"`
	Headers map[string]string      `yaml:"headers"`
	Body    map[string]interface{} `yaml:"body"`
}
