package global

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"test-package-02/app/global/structs"
)

var Config *structs.DBConfig

func Start() (err error) {
	env := os.Getenv("ENV")

	envPathList := []string{
		"env/" + env + "/db.yaml",
	}

	for _, path := range envPathList {
		configFile, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err.Error())
		}

		if err = yaml.Unmarshal(configFile, &Config); err != nil {
			panic(err.Error())
		}
	}

	return nil
}