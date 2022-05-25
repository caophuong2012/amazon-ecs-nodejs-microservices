package controllers

import "net/http"

// HealthCheck func for api
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
