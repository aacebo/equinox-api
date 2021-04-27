package organizations

import (
	"database/sql"
	"time"
)

type Model struct {
	ID        string    `json:"id"`
	Slug      string    `json:"slug"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewModel(rows *sql.Rows) *Model {
	if !rows.Next() {
		return nil
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
		log.Error(err)
	}

	return model
}

func NewModels(rows *sql.Rows) []*Model {
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
			log.Error(err)
		}

		res = append(res, model)
	}

	return res
}
