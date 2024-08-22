package main

import (
	"fmt"
	"log"
	"tank-game/internal/network"

	"tank-game/configs"
	"tank-game/internal/game"
	"tank-game/internal/storage"
)

func main() {
	config, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	dbConfig := storage.DatabaseConfig{
		Host:     config.Database.Host,
		Port:     config.Database.Port,
		User:     config.Database.User,
		Password: config.Database.Password,
		Name:     config.Database.Name,
		SSLMode:  config.Database.SSLMode,
	}

	db, err := storage.NewPostgresDB(dbConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	redisClient, err := storage.NewRedisClient(fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port))
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	defer redisClient.Close()

	gameState := game.NewGameState()

	// Если функция `NewGameState` должна принимать аргументы, убедитесь, что они правильные:
	// gameState := game.NewGameState(db, redisClient)

	// Запуск сервера
	server := network.NewServer(config.Server.Port, gameState)
	log.Printf("Starting server on port %s", config.Server.Port)
	if err := server.Start(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
