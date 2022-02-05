package main

import (
	"github.com/lakshay994/news-aggregator/server"
	"github.com/rs/zerolog"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	server.Serve()
}
