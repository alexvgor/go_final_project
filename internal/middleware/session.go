package session

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/alexvgor/go_final_project/internal/setup"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
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
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			signingMethodErr := fmt.Sprintf("invalid token signing method - %s", token.Header["alg"])
			slog.Error(signingMethodErr)
			return nil, errors.New(signingMethodErr)
		}
		return []byte(session.secret), nil
	})
	if err == nil && token.Valid {
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			return passwordHashValidation(session.password, claims["hash"].(string))
		}
	}
	return false
}

func (session *SessionInstance) PasswordValidation(password string) (string, error) {

	if strings.TrimSpace(password) != session.password {
		return "", errors.New("неверный пароль")
	}

	hash, err := hashPassword(session.password)
	if err != nil {
		slog.Error(fmt.Sprintf("unable to create hash for password - %s", err.Error()))
		return "", errors.New("ошибка создания токена")
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"hash": hash,
	}).SignedString([]byte(session.secret))

	if err != nil {
		slog.Error(fmt.Sprintf("unable to create JWToken - %s", err.Error()))
		return "", errors.New("ошибка создания токена")
	}
	return token, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	return string(bytes), err
}

func passwordHashValidation(password string, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
