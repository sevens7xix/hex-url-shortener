package services

import (
	"github.com/sevens7xix/hex-url-shortener/internal/core/ports"
	utilites "github.com/sevens7xix/hex-url-shortener/pkg/utilities"
)

type Service struct {
	ShortenersRepository ports.ShortenersRepository
	generator            utilites.IdGenerator
}

func NewService(repository ports.ShortenersRepository) *Service {
	return &Service{ShortenersRepository: repository}
}

func (s *Service) Shorten(URL string) (string, error) {

}
