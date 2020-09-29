package auth

import (
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gbrlsnchs/jwt/v3"
)

var bearerError = errors.New("Bearer prefix required")

type JwtPayload struct {
	jwt.Payload
}

type Token struct{}

func NewToken() *Token {
	return &Token{}
}

type TokenInterface interface {
	CreateToken() Token
}

func (tk *Token) CreateToken(userId string) (string, error) {
	var hs = jwt.NewHS256([]byte("dev1986"))
	now := time.Now()

	payload := JwtPayload{
		Payload: jwt.Payload{
			Issuer:         userId,
			Subject:        "dev",
			Audience:       jwt.Audience{"http://localhost:3000"},
			ExpirationTime: jwt.NumericDate(now.Add(3000 * time.Second)),
			NotBefore:      jwt.NumericDate(now),
			IssuedAt:       jwt.NumericDate(now),
			JWTID:          "foobar",
		},
	}

	token, err := jwt.Sign(payload, hs)

	log.Println(err)

	if err != nil {
		return "", err
	}

	return string(token), nil
}

func VerifyToken(r *http.Request) error {
	var secret = jwt.NewHS256([]byte("dev1986"))
	var payload JwtPayload

	now := time.Now()
	token, err := extractToken(r)

	if err != nil {
		return err
	}

	expValidator := jwt.ExpirationTimeValidator(now)
	nbfValidator := jwt.NotBeforeValidator(now)
	validatePayload := jwt.ValidatePayload(&payload.Payload, expValidator, nbfValidator)

	_, err = jwt.Verify([]byte(token), secret, &payload, validatePayload)

	if err != nil {
		return err
	}

	return nil
}

func extractToken(r *http.Request) (string, error) {
	token := r.Header.Get("Authorization")

	bearToken := strings.Split(token, " ")

	if len(bearToken) > 0 && bearToken[0] != "Bearer" {
		return "", bearerError
	}

	if len(bearToken) == 2 {
		return bearToken[1], nil
	}

	return "", nil
}
