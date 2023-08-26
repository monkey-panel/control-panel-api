package config

type LogConfig struct {
	Level     string
	TimeZone  string
	LogName   string
	LogSuffix string
	MaxBackup int
}
