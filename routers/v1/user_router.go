package v1

import (
	"github.com/gin-gonic/gin"
	"go-blog/api"
)

type RouterUser struct{}

func (rs *RouterUser) InitUserRouter(Router *gin.RouterGroup) {
	//userRouter := Router.Group("user").User(middlware.FUnc())
	//userRouterWithoutRecord := Router.Group("user")
	baseApi := api.GroupApi.V1ApiGroup.BaseApi
	{
		//userRouterWithoutRecord.POST("register", baseApi.Register)
		//userRouterWithoutRecord.POST("login", baseApi.Login)
		Router.POST("register", baseApi.Register)
		Router.POST("login", baseApi.Login)
	}

}
