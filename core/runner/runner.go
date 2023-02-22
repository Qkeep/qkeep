package runner

import (
	"fmt"
	"io/ioutil"

	"github.com/Qkeep/qkeep/lib/util"
	"github.com/Qkeep/qkeep/model"
	"gopkg.in/yaml.v3"
)

type Runner struct {
	Globals *model.Global
	Root    string
}

func NewRunner(root string) *Runner {
	return &Runner{
		Globals: &model.Global{},
		Root:    root,
	}
}

func (r *Runner) Run(spawn int) {
	// Load global
	global, err := r.LoadGlobal()
	util.CheckPanic(err)
	r.Globals = global

	// Load global models
	models, err := util.ListAllFiles(r.Root + "/model")
	util.CheckPanic(err)

	var modelPaths []string = util.MapStringArr(models, func(s string) string {
		return r.Root + "/model/" + s
	})

	var globalModels map[string]*model.Model = make(map[string]*model.Model)
	for _, modelPath := range modelPaths {
		modelContent, err := ioutil.ReadFile(modelPath)
		util.CheckPanic(err)

		var model model.Model
		err = yaml.Unmarshal(modelContent, &model)
		util.CheckPanic(err)

		var modelName string
		if model.Name == "" {
			modelName = util.SnakeCaseToUpperCamelCase(util.GetFilenameWithoutExt(modelPath))
		} else {
			modelName = model.Name
		}

		globalModels[modelName] = &model
	}
	r.Globals.Models = globalModels

	// Load all suits
	suitNames, err := util.ListAllFolder(r.Root)
	util.CheckPanic(err)

	suitNames = util.FilterStringArr(suitNames, func(s string) bool {
		return s != "model"
	})
	suitPaths := util.MapStringArr(suitNames, func(s string) string {
		return r.Root + "/" + s
	})

	var suits []*model.Suit
	for _, suitPath := range suitPaths {
		data, err := ioutil.ReadFile(suitPath + "/suit.yaml")
		util.CheckPanic(err)

		var suit model.Suit
		err = yaml.Unmarshal(data, &suit)
		util.CheckPanic(err)

		suit.Path = suitPath

		// Get suit models
		modelNames, err := util.ListAllFiles(suitPath + "/model")
		util.CheckPanic(err)

		var modelPaths []string = util.MapStringArr(modelNames, func(s string) string {
			return suitPath + "/model/" + s
		})

		var suitModels map[string]*model.Model = make(map[string]*model.Model)
		for _, modelPath := range modelPaths {
			suitModelContent, err := ioutil.ReadFile(modelPath)
			util.CheckPanic(err)

			var model model.Model
			err = yaml.Unmarshal(suitModelContent, &model)
			util.CheckPanic(err)

			var modelName string
			if model.Name == "" {
				modelName = util.SnakeCaseToUpperCamelCase(util.GetFilenameWithoutExt(modelPath))
			} else {
				modelName = model.Name
			}

			fmt.Println(model.Name)

			suitModels[modelName] = &model
		}
		suit.Models = suitModels
		suits = append(suits, &suit)
	}

	// Start run suits
	if spawn == 1 {
		// Run suits sequentially
		fmt.Println("Run suits sequentially")
		r.RunSequence(suits)
	} else if spawn == -1 {
		// Run suits concurrently
		fmt.Println("Run suits concurrently")
		fmt.Println(suits)
	} else {
		// Run suits concurrently with limited goroutines
		fmt.Println("Run suits concurrently with limited goroutines")
		fmt.Println(suits)
	}
}

func (r *Runner) RunSequence(suits []*model.Suit) {
	for _, suit := range suits {
		flow := suit.Flow
		models := suit.Models
		fmt.Println(models)
		for _, step := range flow {
			scenarioPath := suit.Path + "/" + step + ".yaml"
			var scenario model.Scenario

			content, err := ioutil.ReadFile(scenarioPath)
			util.CheckPanic(err)

			err = yaml.Unmarshal(content, &scenario)
			util.CheckPanic(err)

			fmt.Println(scenario.Cases)
		}
	}
}

func (r *Runner) LoadGlobal() (*model.Global, error) {
	// Load global
	globalPath := r.Root + "/global.yaml"
	content, err := ioutil.ReadFile(globalPath)
	if err != nil {
		return nil, err
	}

	var global model.Global
	err = yaml.Unmarshal(content, &global)
	if err != nil {
		return nil, err
	}

	return &global, nil
}
