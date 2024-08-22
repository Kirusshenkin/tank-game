package game

import (
	"tank-game/internal/models"
)

// MovePlayer перемещает игрока в указанном направлении
func (g *GameState) MovePlayer(id string, direction models.Direction) bool {
	g.mu.Lock()
	defer g.mu.Unlock()

	player, exists := g.Players[id]
	if !exists {
		return false
	}

	newX, newY := player.X, player.Y
	switch direction {
	case models.Up:
		newY--
	case models.Right:
		newX++
	case models.Down:
		newY++
	case models.Left:
		newX--
	}

	if g.Map.IsValidMove(newX, newY) {
		player.X, player.Y = newX, newY
		player.Direction = direction
		return true
	}

	return false
}
