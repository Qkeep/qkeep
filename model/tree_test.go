package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScenarioPaths_AddTestCasePath(t *testing.T) {
	// Arrange
	scenario := &ScenarioPaths{}

	// Act
	scenario.AddTestCasePath("path/to/test/case")

	// Assert
	assert.Equal(t, []string{"path/to/test/case"}, scenario.GetTestCasePaths())
}

func TestScenarioPaths_GetPath(t *testing.T) {
	// Arrange
	scenario := &ScenarioPaths{Path: "path/to/scenario"}

	// Act
	path := scenario.GetPath()

	// Assert
	assert.Equal(t, "path/to/scenario", path)
}

func TestSuitPaths_AddScenarioPath(t *testing.T) {
	// Arrange
	suit := &SuitPaths{
		Path:          "path/to/suit",
		ScenarioPaths: make(map[string]*ScenarioPaths),
	}

	scenarioPaths := ScenarioPaths{
		Path:          "path/to/scenario",
		TestCasePaths: []string{"/path/to/test/case"},
	}
	// Act
	suit.AddScenarioPath("path/to/scenario", scenarioPaths)

	// Assert
	assert.Equal(t, &scenarioPaths, suit.GetScenarioPaths()["path/to/scenario"])
}

func TestSuitPaths_GetPath(t *testing.T) {
	// Arrange
	suit := &SuitPaths{Path: "path/to/suit"}

	// Act
	path := suit.GetPath()

	// Assert
	assert.Equal(t, "path/to/suit", path)
}

func TestTreePaths_GetRoot(t *testing.T) {
	// Arrange
	tree := &TreePaths{Root: "path/to/root"}

	// Act
	root := tree.GetRoot()

	// Assert
	assert.Equal(t, "path/to/root", root)
}

func TestTreePaths_AddSuitPath(t *testing.T) {
	// Arrange
	tree := &TreePaths{
		Root:      "path/to/root",
		SuitPaths: make(map[string]*SuitPaths),
	}

	suitPaths := SuitPaths{
		Path:          "path/to/suit",
		ScenarioPaths: make(map[string]*ScenarioPaths),
	}
	// Act
	tree.AddSuitPath("path/to/suit", suitPaths)

	// Assert
	assert.Equal(t, &suitPaths, tree.GetSuitPaths()["path/to/suit"])
}

func TestTreePaths_GetSuitPaths(t *testing.T) {
	// Arrange
	tree := &TreePaths{
		Root:      "path/to/root",
		SuitPaths: make(map[string]*SuitPaths),
	}

	suitPaths := SuitPaths{
		Path:          "path/to/suit",
		ScenarioPaths: make(map[string]*ScenarioPaths),
	}
	// Act
	tree.AddSuitPath("path/to/suit", suitPaths)

	// Assert
	assert.Equal(t, &suitPaths, tree.GetSuitPaths()["path/to/suit"])
}
