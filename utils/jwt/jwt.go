package jwt

import (
	"errors"
	"html"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

const (
	jwtSecretKey string = "student_management"
	JwtExpAT     int    = 24
	JwtExpRT     int    = 24
)

func CreateJWTToken(username string, exp int) (string, error) {
	claims := &Claims{
		Username: username,
		Role:     username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(exp)).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ExtractBearerToken(r *http.Request) string {
	authorizationHeader := r.Header.Get("Authorization")
	bearerToken := strings.Split(authorizationHeader, " ")
	return html.EscapeString(bearerToken[1])
}

func ValidateJWTToken(token string) (*Claims, error) {
	t, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey), nil
	})
	if err != nil {
		return nil, err
	}

	if !t.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := t.Claims.(*Claims)
	if !ok {
		return nil, err
	}

	return claims, nil
}
