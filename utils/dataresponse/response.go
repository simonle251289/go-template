package dataresponse

type successRes struct {
	Data   interface{} `json:"data"`
	Paging interface{} `json:"paging,omitempty"`
	Filter interface{} `json:"filter,omitempty"`
}

func NewListSuccessResponse(data, paging, filter interface{}) *successRes {
	return &successRes{Data: data, Paging: paging, Filter: nil}
}

func NewSuccessResponse(data interface{}) *successRes {
	return NewListSuccessResponse(data, nil, nil)
}
