package config_test

import (
	"os"
	"testing"
	"time"

	"github.com/ffrizzo/acme/internal/config"
)

func TestNewDefaultValues(t *testing.T) {
	cfg := config.New()

	if cfg.SimulateExternalCall != (3 * time.Second) {
		t.Error("Default value for 'SimulateExternalCall' should be 3")
	}

	if cfg.CertExpiration != (10 * 24 * time.Hour) {
		t.Error("Default value for 'CertExpiration' should be 24")
	}
}

func TestNewFromEnvVars(t *testing.T) {
	os.Setenv("CERT_EXPIRATION", "1")
	defer os.Unsetenv("CERT_EXPIRATION")

	cfg := config.New()
	if cfg.CertExpiration != (1 * 24 * time.Hour) {
		t.Error("Default value for 'CertExpiration' should be 1")
	}
}
