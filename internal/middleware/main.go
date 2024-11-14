package middleware

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/ngocthanh06/chatapp/internal/models"
	"github.com/ngocthanh06/chatapp/internal/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"os"
	"time"
)

type Claims struct {
	jwt.RegisteredClaims
	Username string `json:"username"`
	Id       string `json:"id"`
}

var claims Claims

func verifyToken(tokenString string) (interface{}, error) {
	// check jwt token access full
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_KEY")), nil
	})

	// check time
	expiredTime := time.Until(claims.ExpiresAt.Time)

	var user models.User

	if err != nil || !token.Valid || expiredTime.Seconds() <= 0 {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid or expired token %v", err)
	}

	result := storage.GetDB().Table("users").Where("id", claims.Id).First(&user)

	if result.RowsAffected == 0 {
		log.Println("User Auth not found")

		return nil, status.Errorf(codes.Unauthenticated, "Invalid or expired token")
	}

	return nil, nil
}
