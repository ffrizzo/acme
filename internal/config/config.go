package config

import (
	"log"
	"os"
	"strconv"
	"time"
)

// Config has the app configurations
type Config struct {
	Port                 string
	SimulateExternalCall time.Duration // seconds
	CertExpiration       time.Duration // days
}

// New creates a config
func New() Config {
	cfg := Config{
		Port:                 ":7070",
		SimulateExternalCall: 3 * time.Second,     //default 3sec
		CertExpiration:       10 * 24 * time.Hour, //default 10days
	}

	if val := os.Getenv("PORT"); len(val) > 0 {
		cfg.Port = val
	}

	if val := os.Getenv("CERT_EXPIRATION"); len(val) > 0 {
		v, err := strconv.Atoi(val)
		if err != nil {
			log.Fatalf("CERT_EXPIRATION is not valid value. %s", err)
		}

		cfg.CertExpiration = time.Duration(v) * 24 * time.Hour
	}

	return cfg
}
