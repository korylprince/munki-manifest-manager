package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/kelseyhightower/envconfig"
)

func main() {
	config := new(Config)
	envconfig.MustProcess("", config)

	m := NewManager(config)
	log.Println("INFO: Listening on:", config.ListenAddr)
	http.ListenAndServe(config.ListenAddr, handlers.CombinedLoggingHandler(os.Stdout, http.StripPrefix(config.StripPrefix, http.HandlerFunc(m.GetManifest))))
}
