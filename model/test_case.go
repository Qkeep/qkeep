package model

type TestCase struct {
	Name  string `yaml:"name"`
	Cases []Case `yaml:"cases"`
}

type Case struct {
	Name       string      `yaml:"name"`
	Request    Request     `yaml:"request"`
	Assertions []Assertion `yaml:"assertions"`
}

type Request struct {
	Path   string                 `yaml:"path"`
	Method string                 `yaml:"method"`
	Body   map[string]interface{} `yaml:"body"`
}

type Assertion struct {
	Name   string      `yaml:"name"`
	On     string      `yaml:"on"`
	Check  string      `yaml:"check"`
	Expect interface{} `yaml:"expect"`
}
