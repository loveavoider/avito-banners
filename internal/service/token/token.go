package token

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

var (
	makeJwtError   = errors.New("")
	errorParseJwt  = errors.New("")
	incorrectToken = errors.New("incorrect token")
)

type service struct {
	secret []byte
}

func (s *service) Generate(aud string) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"aud": aud,
		},
	)

	res, err := token.SignedString(s.secret)

	if err != nil {
		return "", makeJwtError
	}

	return res, nil
}

// TODO метод валидации не в валидаторе

func (s *service) Validate(token string) (*jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	res, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return s.secret, nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}))

	if err != nil {
		return nil, errorParseJwt
	}

	if !res.Valid {
		return nil, incorrectToken
	}

	return &claims, nil
}

func NewTokenService() *service {
	return &service{
		secret: []byte(os.Getenv("JWT_KEY")),
	}
}
