package dto

type Paging struct {
	Limit int `query:"limit" vd:"$>=0" default:"20"`
	Page  int `query:"page" vd:"$>0" default:"1"`
}

func (p *Paging) Offset() int {
	return (p.Page - 1) * p.Limit
}
