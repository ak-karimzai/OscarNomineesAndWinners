package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	ErrInvalidToken = fmt.Errorf("token is not valid")
)

type JWTMaker struct {
	secrectKey string
}

func NewJWTMaker(secret string) Maker {
	return &JWTMaker{
		secrectKey: secret,
	}
}

func (maker *JWTMaker) CreateToken(
	username string, duration time.Duration) (string, error) {
	payload := NewPayload(username, duration)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return jwtToken.SignedString([]byte(maker.secrectKey))
}

func (maker *JWTMaker) VerifyToken(
	token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(maker.secrectKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}
	return payload, nil
}
