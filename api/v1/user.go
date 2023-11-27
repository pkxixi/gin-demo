package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-blog/global"
	"go-blog/models"
	"go-blog/models/request"
	"go-blog/models/response"
	"go-blog/util"
)

type BaseApi struct{}

//
//func LimitFailedCount(c *gin.Context, u *models.User) {
//	v, err := global.Redis.Get(c, u.Username).Result()
//	if err != nil {
//		global.Redis.SetNX(c, u.Username, 1, time.Minute*30)
//	} else {
//		_ := global.Redis.Do(c, "incr", "key", u.Username).Err()
//	}
//}

func (b *BaseApi) Login(c *gin.Context) {
	var requestLogin request.Login
	err := c.ShouldBindJSON(&requestLogin)
	//key := c.ClientIP()
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}
	err = util.Verify(requestLogin, util.LoginVerify)
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}

	// check if user is blocked
	u := &models.User{Username: requestLogin.Username, Password: requestLogin.Password}
	var uBlock = util.NewBlockUtil(u.Username)
	if uBlock.IsBlocked(c) {
		msg := fmt.Sprintf("The number of failed count logins reaches the upper limit, your account is blocked for %d minutes", global.Config.System.LoginLimitTime)
		response.FailWithMsg(msg, c)
		return
	}

	user, e := UserService.Login(u)
	if e != nil {
		global.Logger.Errorf("Login failed, username does not exist or password is wrong, %v", u.Username)
		uBlock.IncrFailedCount(c)
		remainder := uBlock.GetRemainderCount(c)
		if remainder > 0 {
			msg := fmt.Sprintf("Login failed, username or password is wrong, you still have %d chances", remainder)
			response.FailWithMsg(msg, c)
			return
		}
	}
	fmt.Println(user)
	b.TokenNext(c, *user)
	return
}

// TokenNext after login successfully, sign a jwt
func (b *BaseApi) TokenNext(c *gin.Context, user models.User) {
	j := &util.JWT{SignKey: []byte(global.Config.JWT.SignKey)}
	claims := j.CreateClaims(request.BaseClaims{
		UUID:     user.Uuid,
		ID:       user.ID,
		Nickname: user.NickName,
		Username: user.Username,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		global.Logger.Error("Token apply failed")
		response.FailWithMsg("token require failed", c)
		return
	}
	global.Logger.Info(token)
	response.OK(response.LoginResponse{
		User:      user,
		Token:     token,
		ExpiredAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
	}, "login success", c)
	return
}

func (b *BaseApi) Register(c *gin.Context) {
	var requestRegister request.Register
	err := c.ShouldBindJSON(&requestRegister)
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}
	err = util.Verify(requestRegister, util.RegisterVerify)
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}
	user := &models.User{
		Username: requestRegister.Username,
		NickName: requestRegister.NickName,
		Password: requestRegister.Password,
	}
	userReturn, err := UserService.Register(*user)
	if err != nil {
		global.Logger.Errorf("failed to register: %v", err)
		response.Fail(response.UserResponse{User: userReturn}, "failed to register.", c)
		return
	}
	response.OK(response.UserResponse{User: userReturn}, "register success", c)
}
