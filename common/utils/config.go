package utils

import . "github.com/monkey-panel/control-panel-utils/utils"

const globalConfigPath = "data/local_data"

var globalConfig *ConfigStruct

// global configuration struct
type ConfigStruct struct {
	AllowOrigins []string `json:"allow_origins"`
	JWTTimeout   int64    `json:"jwt_timeout"` // hours
	Address      string   `json:"address"`
	EnableTLS    bool     `json:"enable_tls"`
	JWTKey       string   `json:"jwt_key"`
}

// default configuration
func (c ConfigStruct) Default() any {
	return ConfigStruct{
		AllowOrigins: []string{"*"},
		JWTTimeout:   24 * 14, // 14day
		Address:      "0.0.0.0:8000",
		EnableTLS:    true,
		JWTKey:       RandomString(64),
	}
}

func Config() ConfigStruct {
	if globalConfig == nil {
		globalConfig = &ConfigStruct{}
		ConfigRead(globalConfigPath, globalConfig)
	}
	return *globalConfig
}
