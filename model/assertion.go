package model

type Assertion struct {
	Name     string      `yaml:"name"`
	On       string      `yaml:"on"`
	Check    string      `yaml:"check"`
	Expected interface{} `yaml:"expected"`
}
