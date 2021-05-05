package page

import "fmt"

type PageLinks struct {
	First *string `json:"first"`
	Last  *string `json:"last"`
	Next  *string `json:"next"`
	Prev  *string `json:"prev"`
}

func NewLinks(p *Page, total int) *PageLinks {
	var self = new(PageLinks)
	var pages = p.Pages(total)
	var path = p._ctx.Request.URL.Path
	var query = fmt.Sprintf("&perPage=%d", p.PerPage)

	if len(p.Filter) > 0 {
		query = fmt.Sprintf("%s&filter=%s", query, p.Filter)
	}

	var first = fmt.Sprintf("%s?page=1%s", path, query)
	self.First = &first

	if total > 0 {
		var last = fmt.Sprintf("%s?page=%d%s", path, pages, query)
		self.Last = &last
	}

	if p.Page < pages {
		var next = fmt.Sprintf("%s?page=%d%s", path, p.Page+1, query)
		self.Next = &next
	}

	if p.Page > 1 {
		var prev = fmt.Sprintf("%s?page=%d%s", path, p.Page-1, query)
		self.Prev = &prev
	}

	return self
}
