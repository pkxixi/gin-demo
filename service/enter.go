package service

type DemoServiceGroup struct {
	UserServiceGroup UserService
}

var DemoServiceGroupApp = new(DemoServiceGroup)
