package organizations

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"syreclabs.com/go/faker"
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

func Mock() *Model {
	var model = new(Model)
	var name = faker.Company().Name()

	model.ID = uuid.NewString()
	model.Slug = slug.Make(name)
	model.Name = name
	model.CreatedAt = time.Now()
	model.UpdatedAt = time.Now()

	return model
}
