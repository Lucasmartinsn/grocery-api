package Encryptiontoken

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type JWTService_Default struct {
	secretKey string
	issuer    string
}

func NewJWTService_Default() *JWTService_Default {
	return &JWTService_Default{
		secretKey: os.Getenv("JWT_SECRET_KEY_DEFAULT"),
		issuer:    os.Getenv("JWT_ISSUER_DEFAULT"),
	}
}

type Claims_Default struct {
	UserID uuid.UUID `json:"userId"`
	jwt.StandardClaims
}

func (s *JWTService_Default) GenerateToken_Default(userID uuid.UUID) (string, error) {
	claims := &Claims_Default{
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

func (s *JWTService_Default) ValidateToken_Default(tokenString string) (bool, error) {
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
