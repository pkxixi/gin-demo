package models

import "github.com/gofrs/uuid/v5"

// test your code here
type User struct {
	BaseModel
	Uuid     uuid.UUID `json:"uuid" gorm:"index;comment:用户UUID"`
	Username string    `json:"username" gorm:"index;comment:用户登录名"`
	Password string    `json:"-"  gorm:"comment:用户登录密码"`                          // 用户登录密码
	NickName string    `json:"nick-name" gorm:"default:system user;comment:用户昵称"` // 用户昵称
	Phone    string    `json:"phone"  gorm:"comment:用户手机号"`                       // 用户手机号
	Email    string    `json:"email"  gorm:"comment:用户邮箱"`                        // 用户邮箱
	Enable   int       `json:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"`   //用户是否被冻结 1正常 2冻结
	//SideMode    string         `json:"sideMode" gorm:"default:dark;comment:用户侧边主题"`                                          // 用户侧边主题
	//HeaderImg   string         `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"` // 用户头像
	//BaseColor   string         `json:"baseColor" gorm:"default:#fff;comment:基础颜色"`                                           // 基础颜色
	//ActiveColor string         `json:"activeColor" gorm:"default:#1890ff;comment:活跃颜色"`                                      // 活跃颜色
	//AuthorityId uint           `json:"authorityId" gorm:"default:888;comment:用户角色ID"`                                        // 用户角色ID
	//Authority   SysAuthority   `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
	//Authorities []SysAuthority `json:"authorities" gorm:"many2many:sys_user_authority;"`
}

func (User) TableName() string {
	return "gin_demo_users"
}
