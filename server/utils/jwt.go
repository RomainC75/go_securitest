package utils

import (
	db "server/db/sqlc"
	"shared/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func GenerateToken(user db.User) (string, error) {
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"date": jwt.MapClaims{
			"createdAt": now.Unix(),
			"expiresAt": now.Add(time.Second * 2).Unix(),
		},
	})

	jwtSecret := viper.GetString(string(config.SERVER_JWT_SECRET))
	return token.SignedString([]byte(jwtSecret))
}

func ParseToken(tokenString string) (jwt.MapClaims, error) {
	jwtSecret := viper.GetString(string(config.SERVER_JWT_SECRET))
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
