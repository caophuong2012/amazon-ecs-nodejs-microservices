package main

import (
	"os"
	"web/cmd/handlers/routes"
	"web/utils"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"net/http"
)

func main() {

	log.Logger = zerolog.New(os.Stdout).With().Timestamp().Caller().Logger()
	if utils.GetWithDefault("ENV", "DEV") != "PROD" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	r := routes.Route()

	port := utils.GetWithDefault("API_PORT", "8080")
	log.Info().Msg("web api service start on port " + port)
	http.ListenAndServe(":"+port, r)
}
