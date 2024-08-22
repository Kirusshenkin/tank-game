package network

import (
	"fmt"
	"log"
	"net/http"
	"tank-game/internal/handlers"
)

type Server struct {
	port        string
	gameHandler *handlers.GameHandler
}

func NewServer(port string, gameHandler *handlers.GameHandler) *Server {
	return &Server{
		port:        port,
		gameHandler: gameHandler,
	}
}

func (s *Server) Start() {
	// маршрутизация запросов
	SetupRoutes(s.gameHandler)

	log.Printf("Server started at port %s", s.port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", s.port), nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
