package util

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"strings"
	"time"
)

type JWTUtil struct {
}

const (
	ACCESSTOKENTTL  = 5  //60 * 2
	REFRESHTOKENTTL = 10 //60 * 24 * 30
	SECRET          = "6d52e21d599841d0b8c690efa9748ce4"
)

var (
	TokenExpired       = errors.New("token is expired")
	TokenMalformed     = errors.New("token is illegitimate")
	TokenNotValidYet   = errors.New("token not active yet")
	TokenErrorIssuer   = errors.New("wrong issuer")
	TokenErrorAudience = errors.New("token audience is empty")
	TokenInvalid       = errors.New("damaged token")
)

func (ju *JWTUtil) GenerateToken(_id string, sub string, long bool) (string, error) {
	//签发人(issuer)
	//过期时间(expiration time)
	//主题(subject)
	//受众(audience)
	//生效时间(Not Before)
	//签发时间 (Issued At)
	//编号(JWT ID)
	//长/短Token(long)

	jti := uuid.New()
	ts := time.Now()

	var exp time.Time
	if long {
		exp = ts.Add(time.Minute * ACCESSTOKENTTL)
	} else {
		exp = ts.Add(time.Minute * REFRESHTOKENTTL)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:  "melo",
		Subject: sub,
		Audience: jwt.ClaimStrings{
			_id,
		},
		ExpiresAt: &jwt.NumericDate{
			Time: exp,
		},
		NotBefore: &jwt.NumericDate{
			Time: ts,
		},
		IssuedAt: &jwt.NumericDate{
			Time: ts,
		},
		ID: jti.String(),
	})

	tokenS, err := token.SignedString([]byte(SECRET))
	if err != nil {
		return "", err
	}
	return tokenS, nil
}

func (ju *JWTUtil) ParseToken(tokenString string) (string, error) {

	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SECRET), nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return "", TokenExpired
			} else if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return "", TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return "", TokenNotValidYet
			} else if ve.Errors&jwt.ValidationErrorIssuer != 0 {
				return "", TokenErrorIssuer
			} else {
				return "", TokenInvalid
			}
		}
		return "", err
	}

	if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {

		aud := strings.Trim(claims.Audience[0], "")

		if len(aud) <= 0 {
			return "", TokenErrorAudience
		}
		return aud, nil
	}
	return "", TokenInvalid
}
