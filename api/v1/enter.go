package v1

import "go-blog/service"

type ApiV1Group struct {
	BaseApi
}

var (
	UserService = service.DemoServiceGroupApp.UserService
	JWTService  = service.DemoServiceGroupApp.JwtService
)
