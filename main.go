package main

import (
	"log"
	"net/http"
	"tfPortal-gateway/middleware"
	"tfPortal-gateway/routes"
)

func main() {
	router := routes.SetupRoutes()
	log.Println("🚀 API Gateway running at http://localhost:8080")
	http.ListenAndServe(":8080", middleware.WithCORS(router))
}
