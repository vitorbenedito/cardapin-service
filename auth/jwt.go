package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"cardap.in/lambda/apperrors"
	"cardap.in/lambda/model"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

func CreateToken(user model.UserJSON) (string, error) {
	godotenv.Load()
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user"] = user
	claims["exp"] = time.Now().Add(time.Hour * 5).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("api_secret")))
}

func TokenValid(headerValue string) (*model.UserJSON, *apperrors.AppError) {
	godotenv.Load()
	tokenString := extractToken(headerValue)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("api_secret")), nil
	})
	if err != nil || !token.Valid {
		return nil, &apperrors.AppError{err, "Token invalid", http.StatusUnauthorized}
	}
	claims := token.Claims.(jwt.MapClaims)["user"]
	var userJSON model.UserJSON
	if claimsJSON, err := json.Marshal(claims); err != nil {
		return nil, &apperrors.AppError{err, "Token invalid", http.StatusUnauthorized}
	} else {
		if err = json.Unmarshal(claimsJSON, &userJSON); err != nil {
			return nil, &apperrors.AppError{err, "Token invalid", http.StatusUnauthorized}
		}
		return &userJSON, nil
	}
}

func extractToken(bearerToken string) string {
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}
