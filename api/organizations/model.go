package organizations

import (
	"database/sql"
	"time"

	"github.com/aacebo/equinox-api/log"
)

type Model struct {
	ID        string     `json:"id"`
	Key       string     `json:"key"`
	Slug      string     `json:"slug"`
	Name      string     `json:"name"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

func NewModel(rows *sql.Rows) *Model {
	if !rows.Next() {
		return nil
	}

	var model = new(Model)
	var err = rows.Scan(
		&model.ID,
		&model.Key,
		&model.Slug,
		&model.Name,
		&model.CreatedAt,
		&model.UpdatedAt,
	)

	if err != nil {
		log.Error.Fatal(err)
	}

	return model
}

func NewModels(rows *sql.Rows) []*Model {
	var res = []*Model{}

	for rows.Next() {
		var model = new(Model)
		var err = rows.Scan(
			&model.ID,
			&model.Key,
			&model.Slug,
			&model.Name,
			&model.CreatedAt,
			&model.UpdatedAt,
		)

		if err != nil {
			log.Error.Fatal(err)
		}

		res = append(res, model)
	}

	return res
}
