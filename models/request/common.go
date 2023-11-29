package request

type PageInfo struct {
	Page     int    `json:"page" form:"page"`
	PageSize int    `json:"pageSize" form:"pageSize"`
	Keyword  string `json:"keyword" form:"keyword"`
}

type GetById struct {
	ID int `json:"id" form:"id"`
}

func (gbi *GetById) Uint() uint {
	return uint(gbi.ID)
}

type IdsReq struct {
	Ids []int `json:"ids" form:"ids"`
}

type GetAuthorityId struct {
	AuthorityId uint `json:"authorityId" form:"authorityId"`
}

type Empty struct{}
