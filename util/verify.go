package util

var (
	LoginVerify    = Rules{"Username": {NotEmpty()}, "Password": {NotEmpty()}}
	RegisterVerify = Rules{"Username": {NotEmpty()}, "Password": {NotEmpty()}, "NickName": {NotEmpty()}}
)
