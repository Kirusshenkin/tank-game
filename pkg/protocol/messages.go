package protocol

type Message struct {
	Type      string `json:"type"`
	PlayerID  string `json:"player_id"`
	Direction int    `json:"direction,omitempty"`
}
