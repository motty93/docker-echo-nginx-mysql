package model

// Response is api test response struct
type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}
