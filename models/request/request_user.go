package request

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Register struct {
	Username string `json:"username" example:"用户名"`
	Password string `json:"password" example:"密码"`
	NickName string `json:"nickName" example:"昵称"`
	Enable   int    `json:"enable" example:"是否启用"`
	Phone    string `json:"phone" example:"电话号码"`
	Email    string `json:"email" example:"电子邮箱"`
}
