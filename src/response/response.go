package response

import (
	"github.com/aacebo/equinox-api/src/page"
)

type ResponseMeta struct {
	Total int `json:"total"`
	Pages int `json:"pages"`
}

type Response struct {
	Meta  *ResponseMeta   `json:"meta"`
	Links *page.PageLinks `json:"links"`
	Data  interface{}     `json:"data"`
}

func _NewResponseMeta(total int, p *page.Page) *ResponseMeta {
	var m = new(ResponseMeta)

	m.Total = total
	m.Pages = p.Pages(total)

	return m
}

func New(data interface{}) *Response {
	var r = new(Response)

	r.Data = data

	return r
}

func NewPaged(p *page.Page, total int, data interface{}) *Response {
	var r = new(Response)

	r.Meta = _NewResponseMeta(total, p)
	r.Links = p.Links(total)
	r.Data = data

	return r
}
