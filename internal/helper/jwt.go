package helper

import (
	rerror "chatross-api/internal/helper/error"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	access_key = []byte(os.Getenv("ACCESS_KEY"))
	refresh_key = []byte(os.Getenv("REFRESH_KEY"))
) 

type JWTClaims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(userID string) (string, error) {
	claims := JWTClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(access_key)
}

func GenerateRefreshToken(userID string) (string, error) {
	claims := JWTClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(refresh_key)
}

func ValidateAccessToken(tokenString string) (string, error) {
	
	claims := new(JWTClaims)
	token, err := jwt.ParseWithClaims(tokenString, claims, func (token *jwt.Token) (interface{}, error) {
		return access_key, nil
	})

	if err != nil {
		return "", rerror.ErrUnauthorized
	}


	if !token.Valid {
		return "", rerror.ErrUnauthorized
	}

	
	return claims.UserID, nil
}

func ValidateRefreshToken(tokenString string) (string, error) {
	claims := new(JWTClaims)
	token, err := jwt.ParseWithClaims(tokenString, claims, func (token *jwt.Token) (interface{}, error) {
		return refresh_key, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return "", rerror.ErrUnauthorized
		}
		
		return "", rerror.ErrInternalServer
	}


	if !token.Valid {
		return "", rerror.ErrUnauthorized
	}
	
	// if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30 * time.Second {
	// 	return 0, rerror.ErrUnauthorized
	// }

	return claims.UserID, nil
}