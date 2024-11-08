package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

func main() {

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	log.Info().Msg("Welcome to the Track-Tracker repo")

	for {
		log.Info().Msg("Hello all Here!")

		time.Sleep(3 * time.Second)
	}

}
