package datapaging

type Paging struct {
	Page      int   `json:"page" form:"page"`
	Limit     int   `json:"limit" form:"limit"`
	Total     int64 `json:"total" form:"total"`
	TotalPage int   `json:"totalPage"`
	//Support for cursor with UID
	FakeCursor string `json:"cursor" form:"cursor"`
	NextCursor string `json:"next_cursor"`
}

func (p *Paging) Fulfill() {
	if p.Page <= 0 {
		p.Page = 1
	}

	if p.Limit <= 0 {
		p.Limit = 10
	}
}

func (p *Paging) CalculateTotalPage() {
	p.TotalPage = int(p.Total/int64(p.Limit)) + 1
}
