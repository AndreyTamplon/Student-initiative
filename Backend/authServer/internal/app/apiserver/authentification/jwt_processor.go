package authentification

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

type JWTConfig struct {
	JwtSecret string
	TTL       time.Duration
}

const (
	tokenTTL = 1000 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func ConfigureJWT(jwtSecret string, jwtTTL string) (*JWTConfig, error) {
	ttl, err := strconv.Atoi(jwtTTL)
	if err != nil {
		return nil, err
	}
	ttlHours := time.Duration(ttl) * time.Hour
	jwtConfig := &JWTConfig{
		JwtSecret: jwtSecret,
		TTL:       ttlHours,
	}
	return jwtConfig, nil
}

func (config *JWTConfig) GenerateJWT(id int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		id,
	})

	return token.SignedString([]byte(config.JwtSecret))
}

func (config *JWTConfig) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(config.JwtSecret), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}
