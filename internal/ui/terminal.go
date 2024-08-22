package ui

import (
	"fmt"
	"tank-game/internal/models"
)

type TerminalUI struct{}

func NewTerminalUI() *TerminalUI {
	return &TerminalUI{}
}

func (t *TerminalUI) DisplayGameState(state *models.GameState) {
	fmt.Println("Current Game State:")
	fmt.Println("Map:")
	for _, row := range state.Map {
		for _, cell := range row {
			switch cell {
			case 0:
				fmt.Print(". ") // Empty
			case 1:
				fmt.Print("# ") // Wall
			case 2:
				fmt.Print("~ ") // Water
			}
		}
		fmt.Println()
	}

	fmt.Println("\nPlayers:")
	for id, player := range state.Players {
		fmt.Printf("Player %s: Position (%d, %d), Direction: %s\n",
			id, player.X, player.Y, directionToString(player.Direction))
	}
	fmt.Println()
}

func (t *TerminalUI) GetUserInput() (models.Direction, error) {
	var input string
	fmt.Print("Enter direction (W/A/S/D): ")
	_, err := fmt.Scanf("%s", &input)
	if err != nil {
		return 0, err
	}

	switch input {
	case "W", "w":
		return models.Up, nil
	case "D", "d":
		return models.Right, nil
	case "S", "s":
		return models.Down, nil
	case "A", "a":
		return models.Left, nil
	default:
		return 0, fmt.Errorf("invalid direction")
	}
}

func directionToString(d models.Direction) string {
	switch d {
	case models.Up:
		return "Up"
	case models.Right:
		return "Right"
	case models.Down:
		return "Down"
	case models.Left:
		return "Left"
	default:
		return "Unknown"
	}
}
