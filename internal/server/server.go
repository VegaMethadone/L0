package server

import (
	"L0/internal/config"
	"L0/internal/server/handlers"
	"fmt"
	"log"
	"net/http"
	"time"
)

func NewServer() (*http.Server, error) {

	r := handlers.GetAllRoutets()

	conf, err := config.GetConfig()
	if err != nil {
		log.Fatalf("Could not load conf for MUX:  %v", err)
		return nil, err
	}

	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf("%s%s", conf.Network.Address, conf.Network.Port),
		WriteTimeout: time.Second * time.Duration(conf.Network.WriteTimeout),
		ReadTimeout:  time.Second * time.Duration(conf.Network.ReadTimeout),
	}
	return srv, nil
}
