package handlers

import (
	"encoding/json"
	"net/http"
	"tank-game/internal/game"
	"tank-game/internal/models"
)

// MapHandler отвечает за обработку запросов, связанных с картами
type MapHandler struct {
	GameState *game.GameState
}

// NewMapHandler создает новый обработчик для логики карт
func NewMapHandler(gameState *game.GameState) *MapHandler {
	return &MapHandler{GameState: gameState}
}

// LoadMap обрабатывает запрос на загрузку карты
func (h *MapHandler) LoadMap(w http.ResponseWriter, r *http.Request) {
	var gameMap models.Map
	if err := json.NewDecoder(r.Body).Decode(&gameMap); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	h.GameState.Map = game.NewGameMapFromLayout(gameMap.Layout)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Map loaded successfully"))
}

// UpdateMap обрабатывает запрос на обновление карты
func (h *MapHandler) UpdateMap(w http.ResponseWriter, r *http.Request) {
	var updateRequest struct {
		Layout [][]int `json:"layout"`
	}

	if err := json.NewDecoder(r.Body).Decode(&updateRequest); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	h.GameState.Map = game.NewGameMapFromLayout(updateRequest.Layout)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Map updated successfully"))
}
