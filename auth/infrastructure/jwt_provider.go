package infrastructure

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/m-nt/gomod/auth/domain"
)

type JWTProvider struct {
	secret []byte
}

func NewJWTProvider(secret string) *JWTProvider {
	return &JWTProvider{
		secret: []byte(secret),
	}
}

func (j *JWTProvider) Generate(c domain.Claims) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid":   c.UserID,
		"email": c.Email,
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
	})

	return token.SignedString(j.secret)
}

func (j *JWTProvider) Parse(tokenString string) (*domain.Claims, error) {

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return j.secret, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	claims := token.Claims.(jwt.MapClaims)

	return &domain.Claims{
		UserID: int(claims["uid"].(float64)),
		Email:  claims["email"].(string),
	}, nil
}
