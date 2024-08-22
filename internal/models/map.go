package models

// Map представляет игровую карту
type Map struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Layout [][]int `json:"layout"`
}
