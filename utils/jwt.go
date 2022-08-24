package utils

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt"
	"template/modules/auth/authmodels"
	"template/modules/users/usermodels"
	"time"
)

type JwtHash struct {
	Id    string
	Name  string
	Email string
}

func (hash *JwtHash) ToHash() string {
	str := fmt.Sprintf("%s%s%s", hash.Id, hash.Name, hash.Email)
	return HashString(str)
}

func (hash *JwtHash) ToPayload() *map[string]interface{} {
	return &map[string]interface{}{
		"user_id": hash.Id,
	}
}

type JwtWrapper struct {
	SecretKey       string
	Issuer          string
	ExpirationHours int64
}

type JwtClaims struct {
	jwt.StandardClaims `json:",inline"`
	Payload            JwtHash `json:"payload"`
	Hash               string  `json:"hash"`
}

func CreateToken(ttl time.Duration, obj *JwtHash, privateKey string) (string, error) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKey))

	if err != nil {
		return "", fmt.Errorf("create: parse key: %w", err)
	}

	now := time.Now().UTC()

	claims := make(jwt.MapClaims)
	claims["payload"] = obj.ToPayload()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()
	claims["exp"] = now.Add(time.Second * ttl).Unix()
	claims["hash"] = obj.ToHash()

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)

	if err != nil {
		return "", fmt.Errorf("create: sign token: %w", err)
	}

	return token, nil
}

func ValidateToken(token string, publicKey string) *JwtClaims {
	key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKey))
	if err != nil {
		return nil
	}

	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected method: %s", t.Header["alg"])
		}
		return key, nil
	})

	if err != nil {
		return nil
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return nil
	}
	jsonStr, _ := json.Marshal(claims)
	var jwtClaim JwtClaims
	_ = json.Unmarshal(jsonStr, &jwtClaim)
	return &jwtClaim
}

func GenerateToken(accessTTL int, refreshTTL int, rsaPrivateKey string, rsaRefreshPrivateKey string, user *usermodels.UserResponse) authmodels.Token {

	accessToken, _ := CreateToken(time.Duration(accessTTL), &JwtHash{
		Id:    user.ID,
		Name:  fmt.Sprintf("%s %s", user.FirstName, user.LastName),
		Email: user.UserName,
	}, rsaPrivateKey)

	refreshToken, err := CreateToken(time.Duration(refreshTTL), &JwtHash{
		Id:    user.ID,
		Name:  fmt.Sprintf("%s %s", user.FirstName, user.LastName),
		Email: user.UserName,
	}, rsaRefreshPrivateKey)

	if err != nil {
		fmt.Println(err)
	}

	return authmodels.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}
