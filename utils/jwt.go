package utils

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"time"
)

type JwtHash struct {
	Id    string
	Name  string
	Email string
}

type JwtWrapper struct {
	SecretKey       string
	Issuer          string
	ExpirationHours int64
}

type jwtClaims struct {
	jwt.StandardClaims
	Id   string
	Hash string
}

func (w *JwtWrapper) GenerateToken(hashInput JwtHash) (token string, err error) {
	claims := &jwtClaims{
		Id:   hashInput.Id,
		Hash: HashPassword(hashInput.Id + hashInput.Email + hashInput.Name),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(w.ExpirationHours)).Unix(),
			Issuer:    w.Issuer,
		},
	}

	tokenGenerate := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = tokenGenerate.SignedString([]byte(w.SecretKey))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (w *JwtWrapper) ValidateToken(signedToken string) (claim *jwtClaims, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&jwtClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(w.SecretKey), nil
		},
	)

	if err != nil {
		return
	}

	claim, ok := token.Claims.(*jwtClaims)
	if !ok {
		return nil, errors.New("Wrong token")
	}

	if claim.ExpiresAt < time.Now().Local().Unix() {
		return nil, errors.New("Token Expired")
	}
	return claim, nil
}