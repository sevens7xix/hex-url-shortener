package ports

import "github.com/sevens7xix/hex-url-shortener/internal/core/domain"

type ShortenersRepository interface {
	Get(ID uint64) (domain.Data, error)
	Create(Data domain.Data) error
}

type ShortenerService interface {
	Shorten(URL string) (string, error)
	Resolve(ShortURL string) (string, error)
}
