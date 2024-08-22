package network

import (
	"fmt"
	"net/http"

	"tank-game/internal/game"
)

type Server struct {
	port      string
	gameState *game.GameState
}

func NewServer(port string, gameState *game.GameState) *Server {
	return &Server{
		port:      port,
		gameState: gameState,
	}
}

func (s *Server) Start() error {
	http.HandleFunc("/health", s.healthCheck)
	// Другие маршруты

	return http.ListenAndServe(fmt.Sprintf(":%s", s.port), nil)
}

func (s *Server) healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
