package main

import (
	"fmt"
	"identity/cmd/handlers/routes"
	"identity/utils"

	"net/http"
)

func main() {
	r := routes.Route()
	_, err := utils.InitDatabase()
	if err != nil {
		panic(err)
	}
	// TODO : Need to setup logger later
	fmt.Println("init DB success")

	port := utils.GetWithDefault("API_PORT", "3001")
	fmt.Println("listen on port " + port)
	http.ListenAndServe(":"+port, r)
}
