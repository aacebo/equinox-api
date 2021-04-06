package page

import "fmt"

type PageLinks struct {
	First *string `json:"first"`
	Last  *string `json:"last"`
	Next  *string `json:"next"`
	Prev  *string `json:"prev"`
}

func NewLinks(p *Page, total int) *PageLinks {
	var l = new(PageLinks)
	var path = p._ctx.Request.URL.Path
	var query = fmt.Sprintf("&perPage=%d", p.PerPage)

	if len(p.Filter) > 0 {
		query = fmt.Sprintf("%s&filter=%s", query, p.Filter)
	}

	var first = fmt.Sprintf("%s?page=1%s", path, query)

	l.First = &first

	if total > 0 {
		var last = fmt.Sprintf("%s?page=%d%s", path, p.Pages(total), query)
		l.Last = &last
	}

	if p.Page < p.Pages(total) {
		var next = fmt.Sprintf("%s?page=%d%s", path, p.Page+1, query)
		l.Next = &next
	}

	if p.Page > 1 {
		var prev = fmt.Sprintf("%s?page=%d%s", path, p.Page-1, query)
		l.Prev = &prev
	}

	return l
}
