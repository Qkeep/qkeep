package parser

import (
	"github.com/Qkeep/qkeep/model"
	"github.com/Qkeep/qkeep/util"
)

type Parser interface {
	ParseGlobal(filePath string) (*model.Global, error)
	ParseSuit(filePath string) (*model.Suit, error)
	ParseScenario(filePath string) (*model.Scenario, error)
	ParseTestCase(filePath string) (*model.TestCase, error)
	ParseSchema(filePath string) (*model.Schema, error)
	ParseAllSuitePaths(root string) (map[string]*model.SuitPaths, error)
	// ParseTreePaths returns the tree paths of the given root folder.
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

func (p *YamlParser) ParseAllSuitePaths(root string) ([]*model.SuitPaths, error) {
	panic("implement me")
}

// ParseTreePaths returns the tree paths of the given root folder.
func (p *YamlParser) ParseTreePaths(root string) (*model.TreePaths, error) {
	rootAbsPath := util.NewFileUtil().GetAbsolutePathOf(root)

	suitPaths, err := p.ParseAllSuitePaths(rootAbsPath)
	if err != nil {
		return nil, err
	}

	suitPathsMap := make(map[string]*model.SuitPaths)
	for _, suitPath := range suitPaths {
		suitPathsMap[suitPath.Path] = suitPath
	}

	return &model.TreePaths{
		Root:      rootAbsPath,
		SuitPaths: suitPathsMap,
	}, nil
}
