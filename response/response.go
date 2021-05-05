package response

import "github.com/aacebo/equinox-api/page"

type ResponseMeta struct {
	Total int `json:"total"`
	Pages int `json:"pages"`
}

type Response struct {
	Meta  *ResponseMeta   `json:"meta"`
	Links *page.PageLinks `json:"links"`
	Data  interface{}     `json:"data"`
}

func _NewResponseMeta(p *page.Page, total int) *ResponseMeta {
	var self = new(ResponseMeta)

	self.Total = total
	self.Pages = p.Pages(total)

	return self
}

func New(data interface{}) *Response {
	var self = new(Response)

	self.Data = data

	return self
}

func NewPaged(p *page.Page, total int, data interface{}) *Response {
	var self = new(Response)

	self.Meta = _NewResponseMeta(p, total)
	self.Links = p.Links(total)
	self.Data = data

	return self
}
