package models

type LoginResponse struct {
	Status  int `json:"status"`
	Error   string `json:"error"`
	Token   string `json:"token"`
	Timeout string `json:"timeout"`
}