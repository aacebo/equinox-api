package organization

import "time"

type Model struct {
	ID        string    `json:"id" binding:"required"`
	Slug      string    `json:"slug" binding:"required"`
	Name      string    `json:"name" binding:"required"`
	CreatedAt time.Time `json:"createdAt" binding:"required"`
	UpdatedAt time.Time `json:"updatedAt" binding:"required"`
	DeletedAt time.Time `json:"deletedAt"`
}
