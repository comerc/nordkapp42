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

const SECRET = "30b50d8699c8b71ea291f453877e5dec" // TODO: вынести в env

type HasuraClaims struct {
	AllowedRoles []string `json:"x-hasura-allowed-roles"`
	DefaultRole  string   `json:"x-hasura-default-role"`
	UserID       string   `json:"x-hasura-user-id"`
	OrgID        string   `json:"x-hasura-org-id"`
}

type Claims struct {
	jwt.RegisteredClaims
	HasuraClaims HasuraClaims `json:"https://hasura.io/jwt/claims"`
}

func NewAccessToken(claims Claims) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return accessToken.SignedString([]byte(SECRET))
}

func NewRefreshToken(claims jwt.RegisteredClaims) (string, error) {
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return refreshToken.SignedString([]byte(SECRET))
}

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

func ParseAccessToken(accessToken string) (*Payload, error) {
	token, err := jwt.ParseWithClaims(accessToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid JWT token")
	}
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, errors.New("decode JWT token to Claims failed")
	}
	issuedAt, err := claims.GetIssuedAt()
	if err != nil {
		return nil, err
	}
	memberID, err := strconv.Atoi(claims.HasuraClaims.UserID)
	if err != nil {
		return nil, err
	}
	res := Payload{
		MemberID: memberID,
		IssuedAt: issuedAt.Unix(),
	}
	return &res, nil
}

func ParseRefreshToken(refreshToken string) (*jwt.RegisteredClaims, error) {
	token, err := jwt.ParseWithClaims(refreshToken, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid JWT token")
	}
	return token.Claims.(*jwt.RegisteredClaims), err
}

func (payload Payload) IsExpired() bool {
	start := time.Now().Add(-10 * time.Minute).Unix()
	return payload.IssuedAt < start
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
