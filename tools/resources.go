package tools

import (
	"encoding/json"
	"encoding/xml"
	"github.com/mblancoa/go-fun-events/errors"
	"gopkg.in/yaml.v3"
	"io"
	"os"
	"reflect"
	"strconv"
	"time"
)

func LoadJsonConfiguration(fileName string, configObj interface{}) {
	bts := LoadFile(fileName)
	err := json.Unmarshal(bts, configObj)
	errors.ManageErrorPanic(err)
	updateFromEnvironment(reflect.Indirect(reflect.ValueOf(configObj)))
}

func LoadYamlConfiguration(fileName string, configObj interface{}) {
	bts := LoadFile(fileName)
	err := yaml.Unmarshal(bts, configObj)
	errors.ManageErrorPanic(err)
	updateFromEnvironment(reflect.Indirect(reflect.ValueOf(configObj)))
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

func updateFromEnvironment(oValue reflect.Value) {
	oType := oValue.Type()
	for i := 0; i < oType.NumField(); i++ {
		field := oType.Field(i)
		if field.Type.Kind() == reflect.Struct {
			fValue := oValue.Field(i)
			updateFromEnvironment(fValue)
		} else if env, ok := field.Tag.Lookup("env"); ok {
			value := os.Getenv(env)
			if value != "" {
				oValue.Field(i).Set(parseValue(field, value))
			}
		}
	}
}

func parseValue(field reflect.StructField, value string) reflect.Value {
	switch field.Type.Kind() {
	case reflect.String:
		return reflect.ValueOf(value)
	case reflect.Int:
		return parseInt(field, value)
	case reflect.Int64:
		return parseInt64(field, value)
	default:
		return reflect.Value{}
	}
}

func parseInt(field reflect.StructField, value string) reflect.Value {
	intValue, err := strconv.Atoi(value)
	errors.ManageErrorPanic(err)
	return reflect.ValueOf(intValue)
}

func parseInt64(field reflect.StructField, value string) reflect.Value {
	if field.Type.Name() == "Duration" {
		durationValue, err := time.ParseDuration(value)
		errors.ManageErrorPanic(err)
		return reflect.ValueOf(durationValue)
	}
	i64, err := strconv.ParseInt(value, 10, 0)
	errors.ManageErrorPanic(err)
	return reflect.ValueOf(i64)
}
