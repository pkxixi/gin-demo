package models

type SysBaseMenu struct {
	BaseModel
	MenuLevel      uint                   `json:"-"`
	ParentId       string                 `json:"parentId" gorm:"comment:父菜单ID"`     // 父菜单ID
	Path           string                 `json:"path" gorm:"comment:路由path"`        // 路由path
	Name           string                 `json:"name" gorm:"comment:路由name"`        // 路由name
	Hidden         bool                   `json:"hidden" gorm:"comment:是否在列表隐藏"`     // 是否在列表隐藏
	Component      string                 `json:"component" gorm:"comment:对应前端文件路径"` // 对应前端文件路径
	Sort           int                    `json:"sort" gorm:"comment:排序标记"`          // 排序标记
	SysAuthorities []SysAuthority         `json:"authoritys" gorm:"many2many:sys_authority_menus;"`
	Children       []SysBaseMenu          `json:"children" gorm:"-"`
	Parameters     []SysBaseMenuParameter `json:"parameters"`
	//MenuBtn       []SysBaseMenuBtn       `json:"menuBtn"`
	//Meta          `json:"meta" gorm:"embedded;comment:附加属性"`                            // 附加属性

}

type SysBaseMenuParameter struct {
	BaseModel
	SysBaseMenuID uint
	Type          string `json:"type" gorm:"comment:地址栏携带参数为params还是query"` // 地址栏携带参数为params还是query
	Key           string `json:"key" gorm:"comment:地址栏携带参数的key"`            // 地址栏携带参数的key
	Value         string `json:"value" gorm:"comment:地址栏携带参数的值"`            // 地址栏携带参数的值
}

func (SysBaseMenu) TableName() string {
	return "demo_menus"
}
