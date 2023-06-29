package types

import (
	"time"

	"github.com/a3510377/control-panel-api/common/utils"
	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("test")

type (
	JWT    string
	Claims struct {
		ID ID `json:"id"`
		jwt.RegisteredClaims
	}
	RefreshToken struct {
		Token          JWT  `json:"token"`
		ExpirationTime Time `json:"expiration"`
	}
)

func NewJWT(userID ID) (token *RefreshToken) {
	return CreateJWT(Claims{ID: userID}, time.Duration(utils.Config().JWTTimeout))
}

func CreateJWT(claims Claims, newTime time.Duration) (token *RefreshToken) {
	expirationTime := time.Now().Add(newTime)

	claims.RegisteredClaims = jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(expirationTime)}

	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims).SignedString(jwtKey)
	if err != nil {
		return nil
	}

	return &RefreshToken{JWT(tokenString), NewTime(expirationTime)}
}

// JWT to string
func (j JWT) String() string { return string(j) }

func (j JWT) Info() *Claims {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(j.String(), claims, func(token *jwt.Token) (any, error) {
		return jwtKey, nil
	})

	if token == nil || (token != nil && !token.Valid) || err != nil {
		return nil
	}

	return claims
}

// Refresh Token
func (j *JWT) Refresh(newTime time.Duration) *RefreshToken {
	claims := j.Info()
	if claims == nil {
		return nil
	}

	if token := CreateJWT(*claims, newTime); token != nil {
		*j = JWT(token.Token)
		return token
	}

	return nil
}
