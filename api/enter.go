package api

import (
	"go-blog/api/System"
	v1 "go-blog/api/v1"
)

type DemoApiGroup struct {
	SystemApiGroup System.ApiSystemGroup
	V1ApiGroup     v1.ApiV1Group
}

var GroupApi = new(DemoApiGroup)
