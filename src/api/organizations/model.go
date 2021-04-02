package organizations

import (
	"database/sql"
	"time"
)

type Model struct {
	ID        string    `json:"id" binding:"required"`
	Slug      string    `json:"slug" binding:"required"`
	Name      string    `json:"name" binding:"required"`
	CreatedAt time.Time `json:"createdAt" binding:"required"`
	UpdatedAt time.Time `json:"updatedAt" binding:"required"`
	DeletedAt time.Time `json:"deletedAt"`
}

func NewModel(rows *sql.Rows) (*Model, error) {
	if !rows.Next() {
		return nil, nil
	}

	var model = new(Model)
	var err = rows.Scan(
		&model.ID,
		&model.Slug,
		&model.Name,
		&model.CreatedAt,
		&model.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return model, nil
}

func NewModels(rows *sql.Rows) ([]*Model, error) {
	var res = []*Model{}

	for rows.Next() {
		var model = new(Model)
		var err = rows.Scan(
			&model.ID,
			&model.Slug,
			&model.Name,
			&model.CreatedAt,
			&model.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		res = append(res, model)
	}

	return res, nil
}
