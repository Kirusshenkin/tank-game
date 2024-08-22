package handlers

import (
	"encoding/json"
	"net/http"
	"tank-game/internal/game"
	"tank-game/internal/models"
)

// BotHandler отвечает за обработку запросов, связанных с ботами
type BotHandler struct {
	GameState *game.GameState
}

// NewBotHandler создает новый обработчик для логики ботов
func NewBotHandler(gameState *game.GameState) *BotHandler {
	return &BotHandler{GameState: gameState}
}

// AddBot обрабатывает запрос на добавление бота в игру
func (h *BotHandler) AddBot(w http.ResponseWriter, r *http.Request) {
	var bot models.Bot
	if err := json.NewDecoder(r.Body).Decode(&bot); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	h.GameState.AddBot(&bot)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Bot added to the game"))
}

// MoveBot обрабатывает запрос на перемещение бота
func (h *BotHandler) MoveBot(w http.ResponseWriter, r *http.Request) {
	var moveRequest struct {
		ID string `json:"id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&moveRequest); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if h.GameState.MoveBot(moveRequest.ID) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Bot moved successfully"))
	} else {
		http.Error(w, "Invalid move", http.StatusBadRequest)
	}
}
