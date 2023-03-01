package model

type TreePaths struct {
	Root      string
	SuitPaths map[string]SuitPaths
}

type SuitPaths struct {
	Path          string
	ScenarioPaths map[string]ScenarioPaths
}

type ScenarioPaths struct {
	Path          string
	TestCasePaths []string
}

// ScenarioPaths
func (s *ScenarioPaths) AddTestCasePath(path string) {
	s.TestCasePaths = append(s.TestCasePaths, path)
}

func (s *ScenarioPaths) GetTestCasePaths() []string {
	return s.TestCasePaths
}

func (s *ScenarioPaths) GetPath() string {
	return s.Path
}

// SuitPaths
func (s *SuitPaths) AddScenarioPath(path string, scenarioPaths ScenarioPaths) {
	s.ScenarioPaths[path] = scenarioPaths
}

func (s *SuitPaths) GetScenarioPaths() map[string]ScenarioPaths {
	return s.ScenarioPaths
}

func (s *SuitPaths) GetPath() string {
	return s.Path
}

// TreePaths
func (t *TreePaths) AddSuitPath(path string, suitPaths SuitPaths) {
	t.SuitPaths[path] = suitPaths
}

func (t *TreePaths) GetSuitPaths() map[string]SuitPaths {
	return t.SuitPaths
}

func (t *TreePaths) GetRoot() string {
	return t.Root
}
