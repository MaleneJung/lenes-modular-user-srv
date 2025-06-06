package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/MaleneJung/go-mux-frontend-routing-walker/frontend"
	"github.com/MaleneJung/lenes-modular-user-srv/backend"
)

type Config struct {
	Port uint16    `json:"port"`
	TLS  TLSConfig `json:"tls"`
}

type TLSConfig struct {
	Active      bool   `json:"active"`
	Certificate string `json:"certificate"`
	Key         string `json:"key"`
}

func main() {

	configBuffer, err := os.ReadFile("config.json")
	if err != nil {
		log.Fatal(err)
		return
	}

	config := Config{
		Port: 8081,
		TLS: TLSConfig{
			Active:      false,
			Certificate: "tls/certificate.pem",
			Key:         "tls/key.pem",
		},
	}
	if err := json.Unmarshal(configBuffer, &config); err != nil {
		log.Fatal(err)
	}

	serveMux := http.NewServeMux()

	if err := frontend.MuxFrontendWalker(serveMux, "/", "lenes-modular-user/", true); err != nil {
		log.Fatal(err)
		return
	}

	api := backend.NewAPI()
	if err := api.LoadFromFile(); err != nil {
		log.Fatal(err)
		return
	}
	api.RegisterRouteHandlers(serveMux)
	if err := api.SaveToFile(); err != nil {
		log.Fatal(err)
		return
	}

	if config.TLS.Active {
		http.ListenAndServeTLS(":"+fmt.Sprint(config.Port), config.TLS.Certificate, config.TLS.Key, serveMux)
	} else {
		http.ListenAndServe(":"+fmt.Sprint(config.Port), serveMux)
	}

}
