package initial

import (
	"go-blog/config"
	"go-blog/global"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

func InitConf() {
	const ConfigFile = "settings.yaml"
	c := &config.Config{}
	yamlConf, err := os.ReadFile(ConfigFile)
	if err != nil {
		log.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("config initial unmarshal: %v\n", err)
	}
	log.Println("config yaml initial success.")
	//fmt.Println(c)
	global.Config = c
}
