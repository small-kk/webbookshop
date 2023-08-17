package model

type Page struct {
	Books       []*Book //存放当前页图书信息
	PageNo      int     //页码
	PageSize    int     //每页显示记录数
	TotalPageNo int     //总页数，通过计算得到
	TotalRecord int     //总记录数，通过查询数据库得到
	MinPrice    string
	MaxPrice    string
	IsLogin     bool
	UserName    string
}

// IsHasPrev 判断是否有上一页
func (p *Page) IsHasPrev() bool {
	return p.PageNo > 1
}

// IsHasNext 判断是否有下一页
func (p *Page) IsHasNext() bool {
	return p.PageNo < p.TotalPageNo
}

// GetPrevPageNo 获取上一页
func (p *Page) GetPrevPageNo() int {
	if p.IsHasPrev() {
		return p.PageNo - 1
	} else {
		return p.PageNo
	}
}

// GetNextPageNo 获取下一页
func (p *Page) GetNextPageNo() int {
	if p.IsHasNext() {
		return p.PageNo + 1
	} else {
		return p.PageNo
	}
}
