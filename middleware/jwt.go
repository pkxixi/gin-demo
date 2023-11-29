package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go-blog/global"
	"go-blog/models/response"
	"go-blog/service"
	"go-blog/util"
	"strconv"
	"time"
)

var jwtService = service.DemoServiceGroupApp.JwtService

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("x-token")
		if token == "" {
			response.Fail(gin.H{"reload": true}, "you have not logged in or illegal attempt", c)
			c.Abort()
			return
		}
		j := util.NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			if errors.Is(err, util.TokenExpired) {
				response.Fail(gin.H{"reload": true}, "token is expired", c)
				c.Abort()
				return
			}
			response.Fail(gin.H{"reload": true}, err.Error(), c)
			c.Abort()
			return
		}

		if claims.ExpiresAt.Unix()-time.Now().Unix() < claims.BufferTime {
			dr, _ := util.ParserDuration(global.Config.JWT.ExpiredTime)
			claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(dr))
			newToken, _ := j.UpdateTokenFromOld(token, *claims)
			newClaims, _ := j.ParseToken(newToken)
			c.Header("new-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt.Unix(), 10))
		}
		c.Set("claims", claims)
		c.Next()
	}
}
