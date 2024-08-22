package handlers

import (
	"encoding/json"
	"net/http"
	"tank-game/internal/game"
	"tank-game/internal/models"
)

// PlayerHandler отвечает за обработку запросов, связанных с игроками
type PlayerHandler struct {
	GameState *game.GameState
}

// NewPlayerHandler создает новый обработчик для логики игрока
func NewPlayerHandler(gameState *game.GameState) *PlayerHandler {
	return &PlayerHandler{GameState: gameState}
}

// JoinGame обрабатывает запрос на присоединение игрока к игре
func (h *PlayerHandler) JoinGame(w http.ResponseWriter, r *http.Request) {
	var player models.Player
	if err := json.NewDecoder(r.Body).Decode(&player); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	h.GameState.AddPlayer(&player)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Player joined the game"))
}

// UpdatePosition обрабатывает запрос на обновление позиции игрока
func (h *PlayerHandler) UpdatePosition(w http.ResponseWriter, r *http.Request) {
	var updateRequest struct {
		ID        string           `json:"id"`
		Direction models.Direction `json:"direction"`
	}

	if err := json.NewDecoder(r.Body).Decode(&updateRequest); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if h.GameState.MovePlayer(updateRequest.ID, updateRequest.Direction) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Player position updated"))
	} else {
		http.Error(w, "Invalid move", http.StatusBadRequest)
	}
}
