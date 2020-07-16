package Common


type Message struct {
	ActionType string      `json:"actionType"`
	Action     interface{} `json:"action"`
}

type ConnectionStatus struct {
	ClientIds  []uint64 `json:"clientIds"`
	Active     bool     `json:"active"`
	MinPlayers int      `json:"minPlayers"`
	MaxPlayers int      `json:"maxPlayers"`
}