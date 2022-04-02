package jwt

import "github.com/dgrijalva/jwt-go"

type DecodedToken struct {
	Iat    int      `json:"iat"`
	Roles  []string `json:"roles"`
	UserId string   `json:"userId"`
	Email  string   `json:"email"`
	Iss    string   `json:"iss"`
}

func GenerateToken(claims *jwt.Token, secret string) (token string) {
	hmacSecretString := secret
	hmacSecret := []byte(hmacSecretString)
	token, _ = claims.SignedString(hmacSecret)

	return
}
