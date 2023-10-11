package config

type AuthenticationRequest struct {
	AuthMethod    string  `json:"authMethod" binding:"required"`
	Token         *string `json:"token"`
	LocalDateTime string  `json:"date" binding:"required"`
	AppID         string  `json:"appId" binding:"required"`
	UserName      *string `json:"userName"`
	Password      *string `json:"password"`
}

type AuthenticationResult struct {
	IsValid bool
	UserID  string
}

type AuthenticationResponse struct {
	Status             string        `json:"status"`
	PlayerProfileImage string        `json:"playerProfileImage"`
	PublisherPlayerId  string        `json:"publisherPlayerId"`
	PlayerName         string        `json:"playerName"`
	Segments           []string      `json:"segments"`
	Balances           []ItemBalance `json:"balances"`
}
