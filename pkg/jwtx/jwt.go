package jwtx

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

func GenerateJWT(email string, userId int) (string, error) {
	secretKey, err := os.ReadFile("./secret/secret.key")
	if err != nil {
		return "", err
	}

	// Set custom claims
	claims := Claims{
		Email:  email,
		UserID: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(), // Set expiry time to 1 hour
			Issuer:    "your_app",
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with secret key
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ValidateJWT(tokenString string, secretKey []byte) (*Claims, error) {
	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return secretKey, nil
	})

	// Check if the token is valid
	if err != nil {
		return nil, err
	}

	// Check if the token is valid and return claims
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
