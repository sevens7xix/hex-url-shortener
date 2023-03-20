package utilities

import (
	"math/rand"
)

var (
	letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
)

type Shortener interface {
	ShortenURL(URL string) string
}

type ShortenerImplementation struct{}

func (s ShortenerImplementation) ShortenURL(URL string) string {
	var shortURL string

	for i := 0; i < 7; i++ {
		shortURL += string(letters[rand.Intn(len(letters))])
	}

	return shortURL
}
