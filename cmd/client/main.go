// cmd/client/main.go
package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gdamore/tcell/v2"
)

type GameState struct {
	Players map[string]Player `json:"players"`
	Map     [][]int           `json:"map"`
}

type Player struct {
	ID        string `json:"id"`
	X         int    `json:"x"`
	Y         int    `json:"y"`
	Direction int    `json:"direction"`
}

var (
	screen     tcell.Screen
	playerID   string
	serverURL  string
	gameStates = make(chan GameState, 100)
)

func main() {
	// Настройка логирования
	logFile, err := os.Create("client.log")
	if err != nil {
		fmt.Printf("Failed to create log file: %v\n", err)
		return
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	// Получение URL сервера из переменной окружения или использование значения по умолчанию
	serverURL = os.Getenv("SERVER_URL")
	if serverURL == "" {
		serverURL = "http://localhost:8888"
	}
	log.Printf("Using server URL: %s", serverURL)

	// Получение ID игрока
	fmt.Print("Enter your player ID: ")
	reader := bufio.NewReader(os.Stdin)
	playerID, err = reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Failed to read player ID: %v", err)
	}
	playerID = strings.TrimSpace(playerID)
	log.Printf("Player ID: %s", playerID)

	// Инициализация экрана
	screen, err = tcell.NewScreen()
	if err != nil {
		log.Fatalf("Failed to create new screen: %v", err)
	}
	if err := screen.Init(); err != nil {
		log.Fatalf("Failed to initialize screen: %v", err)
	}
	defer screen.Fini()

	// Обработка сигналов для корректного завершения
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-signalChan
		screen.Fini()
		fmt.Println("Game terminated")
		os.Exit(0)
	}()

	// Присоединение к игре
	if err := joinGame(playerID); err != nil {
		log.Fatalf("Failed to join game: %v", err)
	}

	// Запуск получения обновлений
	go receiveUpdates()

	// Запуск отрисовки игры
	go func() {
		log.Println("Starting game state processing")
		for state := range gameStates {
			log.Printf("Processing game state: %+v", state)
			drawGame(state)
			log.Println("Game state drawn")
		}
	}()

	// Основной игровой цикл
	for {
		switch ev := screen.PollEvent().(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEscape:
				return
			case tcell.KeyUp:
				movePlayer("up")
			case tcell.KeyRight:
				movePlayer("right")
			case tcell.KeyDown:
				movePlayer("down")
			case tcell.KeyLeft:
				movePlayer("left")
			}
		case *tcell.EventResize:
			screen.Sync()
		}
	}
}

func joinGame(playerID string) error {
	payload := map[string]string{"playerID": playerID}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal join payload: %v", err)
	}

	resp, err := http.Post(serverURL+"/join", "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return fmt.Errorf("failed to send join request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("join request failed with status: %s", resp.Status)
	}

	log.Println("Successfully joined the game")
	return nil
}

func movePlayer(direction string) {
	payload := map[string]string{"playerID": playerID, "direction": direction}
	jsonPayload, _ := json.Marshal(payload)
	resp, err := http.Post(serverURL+"/move", "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Printf("Failed to move: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Move request failed with status: %s", resp.Status)
	}
}

func receiveUpdates() {
	log.Println("Starting to receive updates")
	for {
		resp, err := http.Get(serverURL + "/events")
		if err != nil {
			log.Printf("Failed to connect to event stream: %v", err)
			time.Sleep(time.Second) // Wait before retrying
			continue
		}
		defer resp.Body.Close()
		log.Println("Connected to event stream")

		scanner := bufio.NewScanner(resp.Body)
		for scanner.Scan() {
			line := scanner.Text()
			log.Printf("Received line: %s", line)
			if len(line) > 6 && line[:6] == "data: " {
				var gameState GameState
				if err := json.Unmarshal([]byte(line[6:]), &gameState); err != nil {
					log.Printf("Failed to unmarshal game state: %v", err)
					continue
				}
				log.Printf("Received game state: %+v", gameState)
				gameStates <- gameState
			}
		}
		if err := scanner.Err(); err != nil {
			log.Printf("Error reading event stream: %v", err)
		}
	}
}

func drawGame(state GameState) {
	log.Printf("Drawing game state. Map size: %dx%d, Players: %d", len(state.Map), len(state.Map[0]), len(state.Players))
	screen.Clear()
	for y, row := range state.Map {
		for x, cell := range row {
			switch cell {
			case 0: // Empty
				drawRune(x, y, ' ')
			case 1: // Wall
				drawRune(x, y, '#')
			case 2: // Water
				drawRune(x, y, '~')
			}
		}
		screen.Show()
		log.Println("Game state drawn and shown")
	}
	for _, player := range state.Players {
		var r rune
		switch player.Direction {
		case 0: // Up
			r = '^'
		case 1: // Right
			r = '>'
		case 2: // Down
			r = 'v'
		case 3: // Left
			r = '<'
		}
		drawRune(player.X, player.Y, r)
	}
	screen.Show()
}

func drawRune(x, y int, r rune) {
	screen.SetContent(x, y, r, nil, tcell.StyleDefault)
}
