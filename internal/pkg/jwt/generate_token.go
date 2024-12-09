package jwt

import (
	"acne-scan-api/internal/model/domain"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateAccessToken(userLogin *domain.Users) (string, error) {

	expireTime := time.Now().Add(time.Hour * 48).Unix()
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"exp":expireTime,
		"username":userLogin.Username, 
		"user_id":userLogin.User_id,
		"role":userLogin.Role,
	})

	token,err:=claims.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err!=nil {
		return "",err
	}

	return token, nil
}
