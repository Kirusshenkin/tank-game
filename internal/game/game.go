package game

import (
	"sync"
	"tank-game/internal/models"
)

// GameState представляет текущее состояние игры
type GameState struct {
	Players map[string]*models.Player
	Bots    map[string]*models.Bot
	Map     *GameMap
	Base    *Base // Добавлено поле Base
	mu      sync.RWMutex
}

// NewGameState создает новое состояние игры
func NewGameState() *GameState {
	return &GameState{
		Players: make(map[string]*models.Player),
		Bots:    make(map[string]*models.Bot),
		Map:     NewGameMap(20, 20),
		Base:    NewBase(), // Инициализация базы
	}
}

// GetPlayer возвращает игрока по его ID
func (g *GameState) GetPlayer(id string) (*models.Player, bool) {
	g.mu.RLock()
	defer g.mu.RUnlock()

	player, exists := g.Players[id]
	return player, exists
}

// AddPlayer добавляет игрока в игру
func (g *GameState) AddPlayer(player *models.Player) {
	g.mu.Lock()
	defer g.mu.Unlock()

	g.Players[player.ID] = player
}

// RemovePlayer удаляет игрока из игры
func (g *GameState) RemovePlayer(id string) {
	g.mu.Lock()
	defer g.mu.Unlock()

	delete(g.Players, id)
}

// GetBot возвращает бота по его ID
func (g *GameState) GetBot(id string) (*models.Bot, bool) {
	g.mu.RLock()
	defer g.mu.RUnlock()

	bot, exists := g.Bots[id]
	return bot, exists
}

// AddBot добавляет бота в игру
func (g *GameState) AddBot(bot *models.Bot) {
	g.mu.Lock()
	defer g.mu.Unlock()

	g.Bots[bot.ID] = bot
}

// RemoveBot удаляет бота из игры
func (g *GameState) RemoveBot(id string) {
	g.mu.Lock()
	defer g.mu.Unlock()

	delete(g.Bots, id)
}
