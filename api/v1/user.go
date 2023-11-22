package v1

import (
	"fmt"
	"go-blog/api"
	"go-blog/global"
	"go-blog/models"
	"go-blog/models/request"
	"go-blog/util/response"
	"go-blog/util"
	"github.com/gin-gonic/gin"
)

type BaseApi struct{}

func (b *BaseApi) Login(c *gin.Context) {
	var requestLogin request.Login
	err := c.ShouldBindJSON(&requestLogin)
	//key := c.ClientIP()
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}
	//err = utils.Verify()
	u := &models.User{Username: requestLogin.Username, Password: requestLogin.Password}
	user, e := api.UserService.Login(u)
	if e != nil {
		global.Logger.Error("Login failed, username does not exist or password is wrong")
		response.FailWithMsg("Login failed, username does not exist or password is wrong", c)
	}
	fmt.Println(user)
	b.TokenNext(c, *user)
	return
}

func (b *BaseApi)TokenNext(c *gin.Context, user models.User){
	j := &util.JWT{SignKey: []byte(global.Config.JWT.SignKey)}
	claims := j.CreateClaims(request.BaseClaims{
		UUID: user.Uuid,
		ID: user.ID,
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
}
