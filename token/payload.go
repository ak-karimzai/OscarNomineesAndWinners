package token

import (
	"fmt"
	"time"
)

var (
	ErrExpiredToken = fmt.Errorf("token has expired")
)

type Payload struct {
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewPayload(
	username string, duration time.Duration) *Payload {
	return &Payload{
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
