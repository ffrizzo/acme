package api

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com/dimfeld/httptreemux"
	"github.com/ffrizzo/acme/internal/cache"
	"github.com/ffrizzo/acme/internal/certificate"
	"github.com/ffrizzo/acme/internal/config"
)

// Result represents response paylod
type Result struct {
	Result interface{} `json:"result"`
}

// Error represents error payload
type Error struct {
	Error string `json:"error"`
}

var (
	cfg config.Config
	cer certificate.Cert
	ca  cache.Cache
	re  *regexp.Regexp
)

// API returns handler of a set of routes
func API(config config.Config, cache cache.Cache, c certificate.Cert) *httptreemux.TreeMux {
	cfg = config
	ca = cache
	cer = c

	var err error
	re, err = regexp.Compile(`^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9])\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\-]*[A-Za-z0-9])$`)
	if err != nil {
		log.Fatalf("Error on comiple regex to validate domains. %s", err)
	}

	router := httptreemux.New()
	router.GET("/cert/:domain", cert)

	return router
}

func cert(w http.ResponseWriter, r *http.Request, params map[string]string) {
	domain := params["domain"]
	if !re.MatchString(domain) {
		response(w, Error{"Invalid domain"}, http.StatusBadRequest)
		return
	}

	//If domain already on cache return this
	if cert, valid := ca.Get(domain); valid {
		response(w, Result{cert}, http.StatusOK)
		return
	}

	cert := cer.New(domain)
	if cfg.SimulateExternalCall > -1 {
		time.Sleep(cfg.SimulateExternalCall)
	}

	response(w, Result{cert}, http.StatusOK)
}

func response(w http.ResponseWriter, data interface{}, code int) {
	if code == http.StatusNoContent {
		w.WriteHeader(code)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Printf("marshalling error: %s", err)
		jsonData = []byte("{}")
	}

	// Send the result back to the client.
	_, err = w.Write(jsonData)
	if err != nil {
		log.Printf("writing response error: %s", err)
	}
}
