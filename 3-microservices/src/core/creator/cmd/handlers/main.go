package main

import (
	"creator/cmd/handlers/routes"
	"creator/internal/databases"
	"creator/internal/databases/storefront"
	"creator/utils"
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

	db, err := utils.InitDatabase()
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed to start the database")
	}

	dbStore := databases.DBStore{}
	dbStore.Storefront = storefront.NewManagement(db)

	r := routes.Route(dbStore)

	port := utils.GetWithDefault("API_PORT", "3001")
	log.Info().Msg("creator service start on port " + port)
	http.ListenAndServe(":"+port, r)
}
