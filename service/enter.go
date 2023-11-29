package service

type DemoServiceGroup struct {
	UserService UserService
	JwtService  JwtService
}

var DemoServiceGroupApp = new(DemoServiceGroup)
