package token

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/loveavoider/avito-banners/merror"
)


type service struct {
	secret []byte
}

func (s *service) Generate(aud string) (string, *merror.MError) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"aud": aud,
		},
	)
	
	res, err := token.SignedString(s.secret)

	if err != nil {
		return "", &merror.MError{Message: "error make jwt"}
	}

	return res, nil
}

func (s *service) Validate(token string) (*jwt.MapClaims, *merror.MError) {
	claims := jwt.MapClaims{}
	res, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return s.secret, nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}))

	if err != nil {
		return nil, &merror.MError{Message: "error validate jwt"}
	}

	if !res.Valid {
		return nil, &merror.MError{Message: "incorrect token"}
	}

	return &claims, nil
}

func NewTokenService() *service {
	return &service{
		secret: []byte(os.Getenv("JWT_KEY")),
	}
}