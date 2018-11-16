package main

import (
	"log"
	"net/http"

	"github.com/ffrizzo/acme/internal/api"
	"github.com/ffrizzo/acme/internal/cache"
	"github.com/ffrizzo/acme/internal/certificate"
	"github.com/ffrizzo/acme/internal/config"
	"github.com/robfig/cron"
)

const (
	domain = "acme.com"
)

func main() {
	cfg := config.New()
	cache := cache.New()
	cert := certificate.New(cfg, cache)
	router := api.API(cfg, cache, cert)

	// Generate internal certificate
	cert.New(domain)
	// Add scheduler to auto renew internal certificate
	c := cron.New()
	c.AddFunc("0 0 0 * * *", func() {
		_, err := cert.Renew(domain)
		if err != nil {
			log.Printf("Error on renew %s domain. %s", domain, err)
		}
	})

	log.Fatal(http.ListenAndServe(":7070", router))
}
