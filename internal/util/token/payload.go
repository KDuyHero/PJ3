package token

import (
	"errors"
	"time"
)

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token is expried")
)

type Payload struct {
	UserId    int64     `json:"userId"`
	Name      string    `json:"name"`
	Role      int8      `json:"role"`
	ExpiredAt time.Time `json:"expiredAt"`
}

func NewPayLoad(uid int64, name string, role int8, duration time.Duration) *Payload {
	return &Payload{
		UserId:    uid,
		Name:      name,
		Role:      role,
		ExpiredAt: time.Now().Add(duration),
	}
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
