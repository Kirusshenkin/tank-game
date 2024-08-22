package models

import (
	"time"
)

// Direction представляет направление движения игрока или бота
type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

// Player представляет игрока в игре
type Player struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
	DeviceID  string    `json:"device_id"`
	X         int       `json:"x"`
	Y         int       `json:"y"`
	Direction Direction `json:"direction"`
}
