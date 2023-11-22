package util

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"go-blog/global"
	"go-blog/models/request"
	"strconv"
	"strings"
	"time"
)

type JWT struct {
	SignKey []byte
}

var (
	TokenExpired     = errors.New("token is expired")
	TokenNotValidYet = errors.New("token is not activated yet")
	TokenMalFormed   = errors.New("token's format is wrong")
	TokenInvalid     = errors.New("can not handle this token")
)

//func NewJWT() *JWT {
//	return &JWT{
//		[]byte(global.Config.JWT.SignKey),
//	}
//}

func ParserDuration(d string) (time.Duration, error) {
	d = strings.TrimSpace(d)
	dr, err := time.ParseDuration(d)
	if err == nil {
		return dr, err
	}
	if strings.Contains(d, "d") {
		index := strings.Index(d, "d")
		hour, _ := strconv.Atoi(d[:index])
		dr = time.Hour * 24 * time.Duration(hour)
		ndr, err := time.ParseDuration(d[index+1:])
		if err != nil {
			return dr, nil
		}
		return dr + ndr, nil
	}
	dv, err := strconv.ParseInt(d, 10, 64)
	if err != nil {
		return 0, errors.New("can not parse time duration")
	}
	return time.Duration(dv), err
}

func (j *JWT) CreateClaims(baseClaims request.BaseClaims) request.CustomClaims {
	bufferDuration, _ := ParserDuration(global.Config.JWT.BufferTime)
	expireDuration, _ := ParserDuration(global.Config.JWT.ExpiredTime)
	claims := request.CustomClaims{
		BaseClaims: baseClaims,
		BufferTime: int64(bufferDuration / time.Second),
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  jwt.ClaimStrings{"gin-demo"},
			NotBefore: jwt.NewNumericDate(time.Now().Add(-1000)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireDuration)),
			Issuer:    global.Config.JWT.IssUer,
		},
	}
	return claims
}

func (j *JWT) CreateToken(claims request.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SignKey)
}

// UpdateTokenFromOld Create new token from old token
func (j *JWT) UpdateTokenFromOld(oldToken string, claims request.CustomClaims) (string, error) {
	v, err, _ := global.ConcurrencyControl.Do("JWT:"+oldToken, func() (interface{}, error) {
		return j.CreateToken(claims)
	})
	return v.(string), err
}

func (j *JWT) ParseToken(t string) (*request.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(t, &request.CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SignKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalFormed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*request.CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	} else {
		return nil, TokenInvalid
	}
}
