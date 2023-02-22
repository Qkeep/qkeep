package model

type Global struct {
	BaseUrl string            `yaml:"base_url"`
	Headers map[string]string `yaml:"headers"`
	Models  map[string]*Model
}
