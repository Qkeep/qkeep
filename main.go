package main

import (
	"fmt"
	"io/ioutil"

	"github.com/Qkeep/qkeep/core/runner"
	"github.com/Qkeep/qkeep/model"
	"github.com/go-resty/resty/v2"
	"gopkg.in/yaml.v3"
)

func main() {
	suitPath := "./test/activity/suit.yaml"
	suitContent, err := ioutil.ReadFile(suitPath)
	if err != nil {
		fmt.Println(err)
	}

	var suit model.Suit
	err = yaml.Unmarshal(suitContent, &suit)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(suit)

	qkeepRunner := runner.NewRunner("./test")
	qkeepRunner.Run(1)

	client := resty.New()
	request := client.R()

	res, err := request.Get("https://www.boredapi.com/api/activity")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
		fmt.Println(res.Request.TraceInfo())
	}
}
