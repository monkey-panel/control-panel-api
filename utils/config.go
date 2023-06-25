package utils

import (
	"os"
	"reflect"
)

type Config struct{ path string }

func NewConfig(path string) *Config { return &Config{path} }

func (c Config) Read() {
	os.ReadFile(c.path)
}

func (c Config) Write() {
	os.WriteFile(c.path, []byte{}, 0o666)
}

func mergeObj(old, new map[string]any) map[string]any {
	for k, value := range new {
		if oldValue, ok := old[k]; ok {
			if v, ok := value.(map[string]any); ok {
				// if oldValue is dict
				if v2, ok := oldValue.(map[string]any); ok {
					old[k] = mergeObj(v2, v)
					continue
				}
			} else if v, ok := value.([]any); ok {
				// if oldValue is list
				if _, ok := oldValue.([]any); ok {
					continue
				}
			} else if reflect.TypeOf(oldValue) == reflect.TypeOf(v) {
				continue
			}
		}
		old[k] = value
	}

	return old
}

type ConfigStruct struct {
	AllowOrigins []string `json:"allow_origins"`
}

func (c ConfigStruct) Default() ConfigStruct {
	return ConfigStruct{
		AllowOrigins: []string{"*"},
	}
}
