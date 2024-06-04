package Encryptiontoken

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type JWTService_Customer struct {
	secretKey string
	issuer    string
}

func NewJWTService_Customer() *JWTService_Customer {
	return &JWTService_Customer{
		secretKey: os.Getenv("JWT_SECRET_KEY_CUSTOMER"),
		issuer:    os.Getenv("JWT_ISSUER_CUSTOMER"),
	}
}

type Claims_Customer struct {
	UserID uuid.UUID `json:"userId"`
	jwt.StandardClaims
}

func (s *JWTService_Customer) GenerateToken_Customer(userID uuid.UUID) (string, error) {
	claims := &Claims_Customer{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 3).Unix(), // Token expiration time
			Issuer:    s.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *JWTService_Customer) ValidateToken_Customer(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("invalid token: %v", tokenString)
		}
		return []byte(s.secretKey), nil
	})

	if err != nil {
		return false, err
	}

	return token.Valid, nil
}
