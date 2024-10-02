package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	session "github.com/alexvgor/go_final_project/internal/middleware"
	"github.com/alexvgor/go_final_project/internal/models"
)

type AuthHandler struct {
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (h *AuthHandler) SignIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var signInDTO models.SignInRequest
		err := json.NewDecoder(r.Body).Decode(&signInDTO)
		if err != nil {
			RespondErrorUnableToSignIn(w, errors.New("ошибка десериализации JSON запроса с данными пользователя"))
			return
		}

		token, err := session.Session.PasswordValidation(signInDTO.Password)
		if err != nil {
			RespondErrorUnableToSignIn(w, err)
			return
		}

		Respond(w, models.Response{Token: token})
	}
}
