package service

import (
	"context"
	"go-blog/global"
	"go-blog/util"
)

type JwtService struct{}

func (js *JwtService) GetRedisJWT(userName string) (redisJWT string, err error) {
	redisJWT, err = global.Redis.Get(context.Background(), userName).Result()
	return redisJWT, err
}

func (js *JwtService) SetRedisJWT(jwt string, userName string) (err error) {
	dr, err := util.ParserDuration(global.Config.JWT.ExpiredTime)
	if err != nil {
		return err
	}
	timer := dr
	err = global.Redis.Set(context.Background(), userName, jwt, timer).Err()
	return err
}

//func
