package main

import (
	"identity/cmd/handlers/routes"
	"identity/internal/databases"
	"identity/internal/databases/creator"
	"identity/internal/databases/user"
	"identity/utils"
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
	dbStore.User = user.NewManagement(db)
	dbStore.Creator = creator.NewManagement(db)

	r := routes.Route(dbStore)

	port := utils.GetWithDefault("API_PORT", "3000")
	log.Info().Msg("Identity service start on port " + port)
	http.ListenAndServe(":"+port, r)
}
