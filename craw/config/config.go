package config

import (
	"io/ioutil"
	"github.com/ghodss/yaml"
	"log"
	"fmt"
	"os"
)

func configPath() string {
	return GoRoot() + "/src/sprider/craw/config" + "/config.yaml"
}
func GoRoot() string{
	return os.Getenv("GOPATH")
}
var Config map[interface{}]interface{}

func InitConfig() {
	absPath := configPath()
	fmt.Println(absPath)
	buffer, err := ioutil.ReadFile(absPath)
	err = yaml.Unmarshal(buffer, &Config)
	fmt.Println(Config)
	if err != nil {
		log.Printf("init config err: %v",err)
	}
}