package game

import (
	"math/rand"
	"tank-game/internal/models"
	"time"
)

type BotAction struct {
	Move      bool
	Direction models.Direction
}

func (g *GameState) MoveBot(id string) bool {
	g.mu.Lock()
	defer g.mu.Unlock()

	bot, exists := g.Bots[id]
	if !exists {
		return false
	}

	rand.Seed(time.Now().UnixNano())
	direction := models.Direction(rand.Intn(4))

	newX, newY := bot.X, bot.Y
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
		bot.X, bot.Y = newX, newY
		bot.Direction = direction
		return true
	}

	return false
}

func (g *GameState) ExecuteBotAction(id string, action BotAction) bool {
	if action.Move {
		return g.MoveBot(id)
	}

	// Здесь нужно добавить логику атаки
	return false
}
