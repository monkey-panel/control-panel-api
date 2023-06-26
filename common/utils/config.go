package utils

import (
	"encoding/json"
	"os"
	"reflect"
)

// global configuration struct
type ConfigStruct struct {
	AllowOrigins []string `json:"allow_origins"`
}

// default configuration
func (c ConfigStruct) Default() any {
	return ConfigStruct{
		AllowOrigins: []string{"*"},
	}
}

// global configuration
func Config() ConfigStruct {
	var config ConfigStruct
	if err := Read("config.json", &config); err != nil {
		panic(err)
	}
	return config
}

type BaseConfig interface {
	Default() any
}

// Read config, if file not exist, write default config
func Read(path string, v BaseConfig) error {
	file, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		if err = Write(path, v); err != nil {
			return err
		}
		return Read(path, v)
	}
	if err != nil {
		return err
	}
	return json.Unmarshal(file, &v)
}

// Write config file
func Write(path string, v BaseConfig) error {
	var data any = v
	if _, err := os.Stat(path); os.IsNotExist(err) {
		data = v.Default()
	}
	rb, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return os.WriteFile(path, rb, 0o666)
}

// merge two object
func mergeObj(old, new map[string]any) map[string]any {
	for k, value := range new {
		if oldValue, ok := old[k]; ok {
			oldValueType := reflect.TypeOf(oldValue).Kind()
			valueType := reflect.TypeOf(value).Kind()

			if valueType == reflect.Map && oldValueType == reflect.Map {
				// if is dict
				old[k] = mergeObj(oldValue.(map[string]any), value.(map[string]any))
				continue
			} else if valueType == reflect.Slice && oldValueType == reflect.Slice {
				// if is list
				continue
			} else if oldValueType == valueType {
				continue
			}
		}
		old[k] = value
	}

	return old
}
