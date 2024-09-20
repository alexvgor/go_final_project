package models

type Response struct {
	Id    int64  `json:"id,omitempty"`
	Error string `json:"error,omitempty"`
}
