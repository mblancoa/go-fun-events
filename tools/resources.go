package tools

import (
	"encoding/json"
	"encoding/xml"
	"github.com/mblancoa/go-fun-events/errors"
	"gopkg.in/yaml.v3"
	"io"
	"os"
)

func LoadJsonConfiguration(fileName string, configObj interface{}) {
	bts := LoadFile(fileName)
	err := json.Unmarshal(bts, configObj)
	errors.ManageErrorPanic(err)
}

func LoadYamlConfiguration(fileName string, configObj interface{}) {
	bts := LoadFile(fileName)
	err := yaml.Unmarshal(bts, configObj)
	errors.ManageErrorPanic(err)
}

func LoadFile(file string) []byte {
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

func UnmarshalXmlResource(file string, v interface{}) error {
	bts := LoadFile(file)
	return xml.Unmarshal(bts, v)
}
