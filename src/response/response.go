package response

import (
	"math"

	"github.com/aacebo/equinox-api/src/page"
)

type ResponseMeta struct {
	Total int `json:"total"`
	Pages int `json:"pages"`
}

type ResponseLinks struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Next  string `json:"next"`
	Prev  string `json:"prev"`
}

type Response struct {
	Meta  *ResponseMeta  `json:"meta"`
	Links *ResponseLinks `json:"links"`
	Data  interface{}    `json:"data"`
}

func _NewResponseMeta(total int, p *page.Page) *ResponseMeta {
	var m = new(ResponseMeta)

	m.Total = total
	m.Pages = int(math.Ceil(float64(total) / float64(p.PerPage)))

	return m
}

func _NewResponseLinks(total int, data interface{}) *ResponseLinks {
	var l = new(ResponseLinks)

	return l
}

func New(data interface{}) *Response {
	var r = new(Response)

	r.Data = data

	return r
}

func NewPaged(p *page.Page, total int, data interface{}) *Response {
	var r = new(Response)

	r.Meta = _NewResponseMeta(total, p)
	r.Links = _NewResponseLinks(total, data)
	r.Data = data

	return r
}
