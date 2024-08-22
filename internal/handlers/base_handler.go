package handlers

import (
	"encoding/json"
	"net/http"
	"tank-game/internal/game"
)

// BaseHandler отвечает за обработку запросов, связанных с базами
type BaseHandler struct {
	GameState *game.GameState
}

// NewBaseHandler создает новый обработчик для логики баз
func NewBaseHandler(gameState *game.GameState) *BaseHandler {
	return &BaseHandler{GameState: gameState}
}

// AttackBase обрабатывает запрос на атаку базы
func (h *BaseHandler) AttackBase(w http.ResponseWriter, r *http.Request) {
	var attackRequest struct {
		Damage int `json:"damage"`
	}

	if err := json.NewDecoder(r.Body).Decode(&attackRequest); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	h.GameState.Base.AttackBase(attackRequest.Damage)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Base attacked successfully"))
}

// GetBaseStatus обрабатывает запрос на получение статуса базы
func (h *BaseHandler) GetBaseStatus(w http.ResponseWriter, r *http.Request) {
	status := struct {
		Health int `json:"health"`
	}{
		Health: h.GameState.Base.Health,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(status)
}
