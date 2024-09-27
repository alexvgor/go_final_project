package models

type Response struct {
	Id    int64  `json:"id,omitempty"`
	Error string `json:"error,omitempty"`
	Token string `json:"token,omitempty"`
}

type SignInRequest struct {
	Password string `json:"password"`
}
