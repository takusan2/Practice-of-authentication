package dto

type Token struct {
	Iss string `json:"iss"`
	Sub string `json:"sub"`
	Exp int64  `json:"exp"`
}

type TokenCreate struct {
	UserID string `json:"user_id"`
	Exp    int64  `json:"exp"`
}
