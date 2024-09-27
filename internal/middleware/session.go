package session

import (
	"errors"
	"fmt"
	"hash/fnv"
	"log/slog"
	"net/http"
	"strings"

	"github.com/alexvgor/go_final_project/internal/setup"
	"github.com/golang-jwt/jwt/v5"
)

type SessionInstance struct {
	password string
	secret   string
}

var Session *SessionInstance

func Init() {
	Session = NewSession()
	slog.Info("session middleware was inited")
}

func NewSession() *SessionInstance {
	return &SessionInstance{password: setup.GetSessionPassword(), secret: setup.GetSessionSecret()}
}

func (session *SessionInstance) AuthenticationIsEnabled() bool {
	return len(session.password) > 0
}

func (session *SessionInstance) Validate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if session.AuthenticationIsEnabled() {
			var token string
			cookie, err := r.Cookie("token")
			if err == nil {
				token = cookie.Value
			}

			if !session.TokenValidation(token) {
				http.Error(w, "Authentification required", http.StatusUnauthorized)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}

func (session *SessionInstance) TokenValidation(tokenString string) bool {
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			c_err := fmt.Sprintf("Invalid token signing method - %s", token.Header["alg"])
			slog.Error(c_err)
			return nil, errors.New(c_err)
		}
		return []byte(session.secret), nil
	})
	return err == nil
}

func (session *SessionInstance) PasswordValidation(password string) (string, error) {

	if strings.TrimSpace(password) != session.password {
		return "", errors.New("неверный пароль")
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"hash": hash(password),
	}).SignedString([]byte(session.secret))

	if err != nil {
		return "", errors.New("ошибка создания токена")
	}
	return token, nil
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}
