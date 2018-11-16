package certificate_test

import (
	"testing"

	"github.com/ffrizzo/acme/internal/cache"
	"github.com/ffrizzo/acme/internal/certificate"
	"github.com/ffrizzo/acme/internal/config"
)

func TestNewCertificate(t *testing.T) {
	cfg := config.New()
	cache := cache.New()
	cert := certificate.New(cfg, cache)

	certificate := cert.New("www.test.com")

	if len(certificate.ID) == 0 {
		t.Error("Certificate does not contain generated ID")
	}
}

func TestNewCertificateMultipleRequest(t *testing.T) {
	cfg := config.New()
	cache := cache.New()
	cert := certificate.New(cfg, cache)

	certificate := cert.New("www.test.com")
	if len(certificate.ID) == 0 {
		t.Error("Certificate does not contain generated ID")
	}

	certificate1 := cert.New("www.test.com")
	if len(certificate.ID) == 0 {
		t.Error("Certificate does not contain generated ID")
	}

	if certificate.ID != certificate1.ID {
		t.Error("Certificate generate with different values for same request")
	}
}
