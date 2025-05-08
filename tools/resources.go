package tools

import (
	"encoding/json"
	"github.com/mblanco/Go-Acme-events/errors"
	"gopkg.in/yaml.v3"
	"io"
	"os"
)

func LoadJsonConfiguration(fileName string, configObj interface{}) {
	bts := loadFile(fileName)
	err := json.Unmarshal(bts, configObj)
	errors.ManageErrorPanic(err)
}

func LoadYamlConfiguration(fileName string, configObj interface{}) {
	bts := loadFile(fileName)
	err := yaml.Unmarshal(bts, configObj)
	errors.ManageErrorPanic(err)
}

func loadFile(file string) []byte {
	confFile, err := os.Open(file)
	errors.ManageErrorPanic(err)
	defer func() {
		err := confFile.Close()
		errors.ManageErrorPanic(err)
	}()

	bts, err := io.ReadAll(confFile)
	errors.ManageErrorPanic(err)
	return bts
}
