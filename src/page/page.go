package page

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Page struct {
	Page    int    `form:"page"`
	PerPage int    `form:"perPage"`
	Filter  string `form:"filter"`
}

func New(ctx *gin.Context) *Page {
	var p = new(Page)

	p.Page = 1
	p.PerPage = 10

	ctx.BindQuery(&p)

	return p
}

func (p *Page) Like() string {
	return fmt.Sprintf("%%%s%%", p.Filter)
}

func (p *Page) Skip() int {
	return (p.Page - 1) * p.PerPage
}
