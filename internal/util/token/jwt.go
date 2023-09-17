package token

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var SECRET_KEY = []byte("ThisIsSecretKey")

type jwtMaker struct {
	secretKey     string
	tokenDuration time.Duration
}

func NewJwtMaker(key string, tokenDuration time.Duration) TokenMaker {
	return &jwtMaker{
		secretKey:     key,
		tokenDuration: tokenDuration,
	}
}

func (maker *jwtMaker) GenerateToken(uid int64, name string, role int8) (string, error) {
	payload := NewPayLoad(uid, name, role, maker.tokenDuration)
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenString, err := jwtToken.SignedString([]byte(maker.secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (maker *jwtMaker) VerifyToken(tokenString string) (*Payload, error) {
	keyFunc := func(jwtToken *jwt.Token) (interface{}, error) {
		_, ok := jwtToken.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}

		return []byte(maker.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(tokenString, &Payload{}, keyFunc)
	if err != nil {
		validationErr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(validationErr.Inner, ErrExpiredToken) {
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
