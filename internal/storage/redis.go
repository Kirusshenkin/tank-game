package storage

import (
	"context"
	"encoding/json"
	"time"

	"tank-game/internal/models"

	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	*redis.Client
}

func NewRedisClient(addr string) (*RedisClient, error) {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	// Проверка подключения
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return &RedisClient{Client: client}, nil
}

func (r *RedisClient) SaveGameState(gameState *models.GameState) error {
	data, err := json.Marshal(gameState)
	if err != nil {
		return err
	}

	return r.Set(context.Background(), "game_state", data, 0).Err()
}

func (r *RedisClient) Close() error {
	return r.Client.Close()
}

func (r *RedisClient) GetGameState() (*models.GameState, error) {
	data, err := r.Get(context.Background(), "game_state").Bytes()
	if err != nil {
		return nil, err
	}

	var gameState models.GameState
	err = json.Unmarshal(data, &gameState)
	if err != nil {
		return nil, err
	}

	return &gameState, nil
}
