package main

import (
	"fmt"
	"io/ioutil"

	"github.com/valdym/ders8/yamlproject/models"
	"gopkg.in/yaml.v2"
)

/*

 YAML
	-gopkg.in/yaml.v2
	-http://godoc.org/gopkg.in/yaml.v2

*/

// Araştırma Konusu: Secret Management Tools
func main() {
	fileName := "config.yaml"
	source, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(source) //byte value then we need to do mapping

	//v2

	var config models.Config
	yaml.Unmarshal(source, &config)
	if err != nil {
		fmt.Println(err)
	}

	//fmt.Println(config)
	fmt.Println(config.Settings)

	for _, v := range config.Settings {
		fmt.Println("Ayar: " + v)
	}

}
