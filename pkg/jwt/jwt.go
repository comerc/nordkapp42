package jwt

import (
	"context"
	"errors"
	"log"
	"strconv"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

const Secret = "30b50d8699c8b71ea291f453877e5dec" // TODO: вынести в env

func TrimBearer(authValue string) string {
	if len(authValue) > 7 && strings.ToUpper(authValue[0:6]) == "BEARER" {
		return authValue[7:]
	}
	return ""
}

type Payload struct {
	MemberID int
	IssuedAt int64
}

func ParsePayload(raw string) (*Payload, error) {
	token, err := jwt.Parse(raw, func(token *jwt.Token) (any, error) {
		// do not try to validate iat
		// mapClaims := token.Claims.(jwt.MapClaims)
		// delete(mapClaims, "iat")
		return []byte(Secret), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid JWT token")
	}
	mapClaims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("decode JWT token to mapClaims failed")
	}
	issuedAt, err := mapClaims.GetIssuedAt()
	if err != nil {
		return nil, err
	}
	hasuraClaims, ok := mapClaims["https://hasura.io/jwt/claims"].(map[string]any)
	if !ok {
		return nil, errors.New("decode JWT token to hasuraClaims failed")
	}
	memberID, err := strconv.Atoi(hasuraClaims["x-hasura-user-id"].(string))
	if err != nil {
		return nil, err
	}
	res := Payload{
		MemberID: memberID,
		IssuedAt: issuedAt.Unix(),
	}
	return &res, nil
}

func (payload Payload) IsExpired() bool {
	return payload.IssuedAt < time.Now().Add(-10*time.Minute).Unix()
}

func GetPayload(ctx context.Context) *Payload {
	res, ok := ctx.Value("JWTPayload").(*Payload)
	if !ok {
		log.Println("GetPayload is empty!")
	}
	// без ok при отсутствии JWTPayload:
	// "interface conversion: interface {} is nil, not *Payload"
	return res
}
