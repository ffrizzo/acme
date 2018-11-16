package certificate

import (
	"errors"
	"time"

	"github.com/ffrizzo/acme/internal/cache"
	"github.com/ffrizzo/acme/internal/config"
	uuid "github.com/satori/go.uuid"
)

// Cert interface
type Cert interface {
	New(domain string) Certificate
	Renew(domain string) (*Certificate, error)
}

type cert struct {
	cfg   config.Config
	cache cache.Cache
}

// Certificate model
type Certificate struct {
	Domain         string    `json:"domain"`
	ID             string    `json:"uuid"`
	ExpirationDate time.Time `json:"expirationDate"`
}

// New returns a cert provider
func New(config config.Config, cache cache.Cache) Cert {
	return &cert{
		cfg:   config,
		cache: cache,
	}
}

// Build generate new certificate or return certificate
func (c *cert) New(domain string) Certificate {
	value, valid := c.cache.Get(domain)
	if valid {
		return value.(Certificate)
	}

	t := time.Now()
	expiration := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
	expiration = expiration.Add(c.cfg.CertExpiration)

	cert := Certificate{
		ID:             uuid.NewV4().String(),
		Domain:         domain,
		ExpirationDate: expiration,
	}

	c.cache.Set(domain, cert)
	return cert
}

// Renew generate a new certificate for domain
func (c *cert) Renew(domain string) (*Certificate, error) {
	value, valid := c.cache.Get(domain)
	if !valid {
		return nil, errors.New("Certificate not being created yet")
	}

	cert := value.(Certificate)
	difference := time.Now().Sub(cert.ExpirationDate).Hours() / 24
	if difference > 1 { // Renew is required only if expiry in 1 day
		return &cert, nil
	}

	cert = Certificate{
		ID:             uuid.NewV4().String(),
		Domain:         domain,
		ExpirationDate: time.Now().Add(c.cfg.CertExpiration),
	}

	c.cache.Set(domain, cert)
	return &cert, nil
}
