package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sevens7xix/hex-url-shortener/app/internal/core/services"
	"github.com/sevens7xix/hex-url-shortener/app/internal/handlers"
	"github.com/sevens7xix/hex-url-shortener/app/internal/repositories"
	"github.com/sevens7xix/hex-url-shortener/app/pkg/utilities"
)

func main() {
	repository := repositories.NewDynamoDBRepository()
	utilities := utilities.ShortenerImplementation{}
	service := services.NewService(repository, utilities)
	handler := handlers.NewHTTPHandler(service)

	router := httprouter.New()
	router.GET("/api/v1/generate/:original", handler.ShortenURL)
	router.GET("/api/v1/resolve/:shortURL", handler.ResolveURL)

	log.Fatal(http.ListenAndServe(":8080", router))
}
