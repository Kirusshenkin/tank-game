package network

import (
	"encoding/json"
	"net"
	"tank-game/internal/models"
	"tank-game/pkg/protocol"
)

type Client struct {
	conn     net.Conn
	playerID string
}

func NewClient(address, playerID string) (*Client, error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}

	return &Client{
		conn:     conn,
		playerID: playerID,
	}, nil
}

func (c *Client) SendJoin() error {
	msg := protocol.Message{
		Type:     "join",
		PlayerID: c.playerID,
	}
	return json.NewEncoder(c.conn).Encode(msg)
}

func (c *Client) SendMove(direction models.Direction) error {
	msg := protocol.Message{
		Type:      "move",
		PlayerID:  c.playerID,
		Direction: int(direction),
	}
	return json.NewEncoder(c.conn).Encode(msg)
}

func (c *Client) ReceiveGameState() (*models.GameState, error) {
	var gameState models.GameState
	err := json.NewDecoder(c.conn).Decode(&gameState)
	if err != nil {
		return nil, err
	}
	return &gameState, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}
