package types

import (
	"net/http"
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

func NewJWT(userID ID) (token *RefreshToken, status int) {
	return CreateJWT(Claims{ID: userID}, time.Duration(utils.Config().JWTTimeout))
}

func CreateJWT(claims Claims, newTime time.Duration) (token *RefreshToken, status int) {
	expirationTime := time.Now().Add(newTime)

	claims.RegisteredClaims = jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(expirationTime)}

	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims).SignedString(jwtKey)
	if err != nil {
		return nil, http.StatusInternalServerError
	}

	return &RefreshToken{JWT(tokenString), NewTime(expirationTime)}, http.StatusOK
}

// JWT to string
func (j JWT) String() string { return string(j) }

// state: `200` OK
// state: `401` Unauthorized
// data: token info
func (j JWT) Info() (data *Claims, status int) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(j.String(), claims, func(token *jwt.Token) (any, error) {
		return jwtKey, nil
	})

	// err == jwt.ErrSignatureInvalid ||
	if token == nil || (token != nil && !token.Valid) {
		return nil, http.StatusUnauthorized
	} else if err != nil {
		return nil, http.StatusUnauthorized
	}

	return claims, http.StatusOK
}

// Refresh Token
func (j *JWT) Refresh(newTime time.Duration) (refreshToken *RefreshToken, status int) {
	claims, status := j.Info()
	if status != 200 {
		return nil, status
	}

	token, status := CreateJWT(*claims, newTime)
	if status != http.StatusOK {
		return nil, status
	}

	*j = JWT(token.Token)

	return token, status
}
