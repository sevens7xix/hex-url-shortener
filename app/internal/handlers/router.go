package handlers

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sevens7xix/hex-url-shortener/app/internal/core/ports"
)

type HTTPHandler struct {
	service ports.ShortenerService
}

func NewHTTPHandler(shortenerService ports.ShortenerService) *HTTPHandler {
	return &HTTPHandler{
		service: shortenerService,
	}
}

func (handler *HTTPHandler) ShortenURL(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	originalURL := ps.ByName("original")

	shortURL, err := handler.service.Shorten(originalURL)

	if err != nil {
		http.Error(w, "Cannot shorten URL", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, shortURL)
}

func (handler *HTTPHandler) ResolveURL(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	shortURL := ps.ByName("shortURL")

	originalURL, err := handler.service.Resolve(shortURL)

	if err != nil {
		http.Error(w, "Short URL binding not found!", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, originalURL)

}
