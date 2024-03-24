package jwt

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt"
)

const JwtSecret = "30b50d8699c8b71ea291f453877e5dec"

func ValidateJWT(raw string) (int, error) {
	token, err := jwt.Parse(raw, func(token *jwt.Token) (any, error) {
		// do not try to validate iat
		mapClaims := token.Claims.(jwt.MapClaims)
		delete(mapClaims, "iat")
		return []byte(JwtSecret), nil
	})
	if err != nil {
		return -1, err
	}
	if !token.Valid {
		return -1, errors.New("invalid JWT token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return -1, errors.New("decode JWT token map claims failed")
	}
	fmt.Printf("%#v", claims)

	// if claims, ok := token.Claims.(*Claims); ok && token.Valid {
	// 	fmt.Printf("%#v", claims)
	// 	// fmt.Println((*jwt.MapClaims)(claims)["myCustomField"])
	// } else {
	// 	fmt.Println(err)
	// }

	// token, err := jwt.ParseWithClaims(jwtToken, &Claims{}, func(token *jwt.Token) (any, error) {
	// 	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
	// 		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	// 	}
	// 	return JwtSecret, nil
	// })
	// if err != nil {
	// 	return -1, errors.New("JWT parsed failed")
	// }
	// claims, ok := token.Claims.(*Claims)
	// if ok && token.Valid {
	// 	fmt.Printf("%#v", claims)
	// 	return 1, nil
	// }
	// return -1, errors.New("JWT claims failed")
	return 1, nil
}
