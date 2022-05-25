package main

import (
	"collection/cmd/handlers/routes"
	"collection/internal/databases"
	"collection/utils"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"net/http"
)

func main() {

	log.Logger = zerolog.New(os.Stdout).With().Timestamp().Caller().Logger()
	if utils.GetWithDefault("ENV", "DEV") != "PROD" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	_, err := utils.InitDatabase()
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed to start the database")
	}

	dbStore := databases.DBStore{}

	r := routes.Route(dbStore)

	port := utils.GetWithDefault("API_PORT", "3003")
	log.Info().Msg("collection service start on port " + port)
	http.ListenAndServe(":"+port, r)
}
