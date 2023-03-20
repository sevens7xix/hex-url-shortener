package services

import (
	"time"

	"github.com/sevens7xix/hex-url-shortener/internal/core/domain"
	"github.com/sevens7xix/hex-url-shortener/internal/core/ports"
	utilites "github.com/sevens7xix/hex-url-shortener/pkg/utilities"
)

type Service struct {
	ShortenersRepository ports.ShortenersRepository
	generator            utilites.Shortener
}

func NewService(repository ports.ShortenersRepository, generator utilites.Shortener) *Service {
	return &Service{ShortenersRepository: repository,
		generator: generator}
}

func (s *Service) Shorten(URL string) (string, error) {
	ShortURL := s.generator.ShortenURL(URL)

	newURLPair := domain.NewData(URL, ShortURL, time.Now())

	if err := s.ShortenersRepository.Create(newURLPair); err != nil {
		return "", err
	}

	return newURLPair.Short, nil
}

func (s *Service) Resolve(ShortURL string) (string, error) {
	URLPair, err := s.ShortenersRepository.Get(ShortURL)

	if err != nil || URLPair == (domain.Data{}) {
		return "", err
	}

	return URLPair.Original, nil
}
