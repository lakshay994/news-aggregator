package server

import (
	"fmt"
	"net/http"

	"github.com/lakshay994/news-aggregator/server/handlers"
	"github.com/rs/zerolog/log"
)

func init() {
	http.HandleFunc(HEALTH, handlers.Health)
	http.HandleFunc(NEWS, handlers.NewsHandler)
}

func Serve() {
	log.Info().Msg(fmt.Sprintf("Server listening on %s", PORT))
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Fatal().Msg(fmt.Sprintf("Server initialization failed %v", err))
	}
}
