package jwt

import (
	"errors"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt"
)

const JwtSecret = "30b50d8699c8b71ea291f453877e5dec" // TODO: вынести в env

func TrimBearer(authValue string) string {
	if len(authValue) > 7 && strings.ToUpper(authValue[0:6]) == "BEARER" {
		return authValue[7:]
	}
	return ""
}

func ValidateJWT(raw string) (int, error) {
	token, err := jwt.Parse(raw, func(token *jwt.Token) (any, error) {
		// do not try to validate iat
		// mapClaims := token.Claims.(jwt.MapClaims)
		// delete(mapClaims, "iat")
		return []byte(JwtSecret), nil
	})
	if err != nil {
		return 0, err
	}
	if !token.Valid {
		return 0, errors.New("invalid JWT token")
	}
	mapClaims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("decode JWT token to mapClaims failed")
	}
	hasuraClaims, ok := mapClaims["https://hasura.io/jwt/claims"].(map[string]any)
	if !ok {
		return 0, errors.New("decode JWT token to hasuraClaims failed")
	}
	return strconv.Atoi(hasuraClaims["x-hasura-user-id"].(string))
}
