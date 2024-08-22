package network

import (
	"net/http"
	"tank-game/internal/handlers"
)

func SetupRoutes(gameHandler *handlers.GameHandler) {
	http.HandleFunc("/create_game", gameHandler.CreateGame)
	http.HandleFunc("/join_game", gameHandler.JoinGame)
	http.HandleFunc("/update_position", gameHandler.UpdatePosition)
	http.HandleFunc("/end_game", gameHandler.EndGame)
}
