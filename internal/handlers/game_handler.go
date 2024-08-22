package handlers

import (
	"encoding/json"
	"net/http"
	"tank-game/internal/game"
	"tank-game/internal/models"
)

// GameHandler отвечает за обработку запросов, связанных с игрой
type GameHandler struct {
	GameState *game.GameState
}

// NewGameHandler создает новый обработчик для логики игры
func NewGameHandler(gameState *game.GameState) *GameHandler {
	return &GameHandler{GameState: gameState}
}

// CreateGame обрабатывает запрос на создание новой игры
func (h *GameHandler) CreateGame(w http.ResponseWriter, r *http.Request) {
	// Здесь можно добавить логику инициализации новой игры
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Game created"))
}

// JoinGame обрабатывает запрос на присоединение игрока к игре
func (h *GameHandler) JoinGame(w http.ResponseWriter, r *http.Request) {
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
func (h *GameHandler) UpdatePosition(w http.ResponseWriter, r *http.Request) {
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

// EndGame обрабатывает запрос на завершение игры
func (h *GameHandler) EndGame(w http.ResponseWriter, r *http.Request) {
	// Здесь можно добавить логику завершения игры и сохранения состояния
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Game ended"))
}
