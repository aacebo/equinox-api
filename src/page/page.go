package page

import (
	"fmt"
	"math"

	"github.com/gin-gonic/gin"
)

type Page struct {
	Page    int    `form:"page"`
	PerPage int    `form:"perPage"`
	Filter  string `form:"filter"`
	_ctx    *gin.Context
}

func New(ctx *gin.Context) *Page {
	var p = new(Page)

	p.Page = 1
	p.PerPage = 10
	p._ctx = ctx

	ctx.BindQuery(&p)

	if p.Page < 1 {
		p.Page = 1
	}

	if p.PerPage < 10 || p.PerPage > 100 {
		p.PerPage = 10
	}

	return p
}

func (p *Page) Like() string {
	return fmt.Sprintf("%%%s%%", p.Filter)
}

func (p *Page) Skip() int {
	return (p.Page - 1) * p.PerPage
}

func (p *Page) Pages(total int) int {
	return int(math.Ceil(float64(total) / float64(p.PerPage)))
}

func (p *Page) Links(total int) *PageLinks {
	return NewLinks(p, total)
}
