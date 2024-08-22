package network

import (
	"net/http"
	"tank-game/internal/handlers"
	"tank-game/internal/middleware"
)

func SetupRoutes(gameHandler *handlers.GameHandler) {

	mux := http.NewServeMux()

	mux.Handle("/create_game", middleware.LoggingMiddleware(middleware.RecoverMiddleware(http.HandlerFunc(gameHandler.CreateGame))))
	mux.Handle("/join_game", middleware.LoggingMiddleware(middleware.RecoverMiddleware(http.HandlerFunc(gameHandler.JoinGame))))
	mux.Handle("/update_position", middleware.LoggingMiddleware(middleware.RecoverMiddleware(http.HandlerFunc(gameHandler.UpdatePosition))))
	mux.Handle("/end_game", middleware.LoggingMiddleware(middleware.RecoverMiddleware(http.HandlerFunc(gameHandler.EndGame))))

	http.Handle("/", mux)
}
