package page

import (
	"fmt"
	"math"

	"github.com/gin-gonic/gin"
)

type Page struct {
	Page    int    `form:"page" binding:"min=1"`
	PerPage int    `form:"perPage" binding:"min=5,max=100"`
	Filter  string `form:"filter"`
	_ctx    *gin.Context
}

func New(ctx *gin.Context) *Page {
	var self = new(Page)

	self.Page = 1
	self.PerPage = 10

	if err := ctx.BindQuery(self); err != nil {
		return nil
	}

	self._ctx = ctx

	return self
}

func (self *Page) Like() string {
	return fmt.Sprintf("%%%s%%", self.Filter)
}

func (self *Page) Skip() int {
	return (self.Page - 1) * self.PerPage
}

func (self *Page) Pages(total int) int {
	return int(math.Ceil(float64(total) / float64(self.PerPage)))
}

func (self *Page) Links(total int) *PageLinks {
	return NewLinks(self, total)
}
