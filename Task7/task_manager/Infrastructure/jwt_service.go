package Infrastructure

import (
	"task_manager/Domain"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWTToken(user *Domain.User, env *Env) (string, error) {
	// Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"role":    user.Role,
		"exp":     jwt.TimeFunc().Add(time.Hour * 24).Unix(), // Token expiration time
		"iat":     jwt.TimeFunc().Unix(),                     // Issued at time
	})

	// Sign the token with a secret key
	signedToken, err := token.SignedString([]byte(env.AccessTokenSecret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}