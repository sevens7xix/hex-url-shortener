package ports

import "github.com/sevens7xix/hex-url-shortener/app/internal/core/domain"

type ShortenersRepository interface {
	Get(shortURL string) (domain.Data, error)
	Create(Data domain.Data) error
}

type ShortenerService interface {
	Shorten(URL string) (string, error)
	Resolve(ShortURL string) (string, error)
}
