package parser

import "github.com/Qkeep/qkeep/model"

type Parser interface {
	ParseGlobal(filePath string) (*model.Global, error)
	ParseSuit(filePath string) (*model.Suit, error)
	ParseScenario(filePath string) (*model.Scenario, error)
	ParseTestCase(filePath string) (*model.TestCase, error)
	ParseSchema(filePath string) (*model.Schema, error)
	ParseTreePaths(root string) (*model.TreePaths, error)
}

type YamlParser struct{}

func (p *YamlParser) ParseGlobal(filePath string) (*model.Global, error) {
	panic("implement me")
}

func (p *YamlParser) ParseSuit(filePath string) (*model.Suit, error) {
	panic("implement me")
}

func (p *YamlParser) ParseScenario(filePath string) (*model.Scenario, error) {
	panic("implement me")
}

func (p *YamlParser) ParseTestCase(filePath string) (*model.TestCase, error) {
	panic("implement me")
}

func (p *YamlParser) ParseSchema(filePath string) (*model.Schema, error) {
	panic("implement me")
}

func (p *YamlParser) ParseTreePaths(root string) (*model.TreePaths, error) {
	panic("implement me")
}
