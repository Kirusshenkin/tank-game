package models

// Bot представляет бота в игре
type Bot struct {
	ID        string    `json:"id"`
	Level     int       `json:"level"`
	Health    int       `json:"health"`
	X         int       `json:"x"`
	Y         int       `json:"y"`
	Direction Direction `json:"direction"`
}
