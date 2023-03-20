package repositories

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sevens7xix/hex-url-shortener/app/internal/core/domain"
)

type RedisRepository struct {
	client *redis.Client
}

func NewRedisRepository() *RedisRepository {
	return &RedisRepository{
		client: redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		}),
	}
}

func (repository *RedisRepository) Get(shortURL string) (domain.Data, error) {
	originalURL, err := repository.client.Get(context.Background(), shortURL).Result()

	if err != nil {
		return domain.Data{}, err
	}

	return domain.NewData(originalURL, shortURL), nil
}

func (repository *RedisRepository) Create(Data domain.Data) error {

	if err := repository.client.Set(context.Background(), Data.Short, Data.Original, 24*time.Hour).Err(); err != nil {
		return err
	}

	return nil
}
