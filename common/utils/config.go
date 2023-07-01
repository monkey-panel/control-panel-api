package utils

import (
	"encoding/json"
	"os"
	"reflect"
	"time"
)

// global configuration struct
type ConfigStruct struct {
	AllowOrigins []string `json:"allow_origins"`
	JWTTimeout   int64    `json:"jwt_timeout"`
	Address      string   `json:"address"`
	EnableTLS    bool     `json:"enable_tls"`
	JWTKey       string   `json:"jwt_key"`
}

// default configuration
func (c ConfigStruct) Default() any {
	return ConfigStruct{
		AllowOrigins: []string{"*"},
		JWTTimeout:   int64(time.Hour * 24 * 7),
		Address:      "0.0.0.0:8000",
		EnableTLS:    true,
	}
}

// global configuration
func Config() ConfigStruct {
	var config ConfigStruct
	if err := Read("data/local_data", &config); err != nil {
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
