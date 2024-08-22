package models

// GameState представляет текущее состояние игры
type GameState struct {
	Players map[string]Player `json:"players"`
	Bots    map[string]Bot    `json:"bots"`
	Map     [][]int           `json:"map"`
}
