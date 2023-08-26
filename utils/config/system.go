package config

type SystemConfig struct {
	Host   string
	Port   string
	DBFile string

	SSLMode string // disable, require, verify-ca, verify-full
	SSLCert string // path to certificate file
	SSLKey  string // path to private key file
}
