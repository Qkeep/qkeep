package model

type Scenario struct {
	Name  string `yaml:"name"`
	Cases []Case `yaml:"cases"`
}

type Case struct {
	Name       string      `yaml:"name"`
	Request    Request     `yaml:"request"`
	Assertions []Assertion `yaml:"assertions"`
}
