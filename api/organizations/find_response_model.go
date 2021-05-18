package organizations

import (
	"database/sql"
	"time"

	"github.com/aacebo/equinox-api/log"
)

type FindResponseModel struct {
	Slug      string     `json:"slug"`
	Name      string     `json:"name"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

func NewFindResponseModel(rows *sql.Rows) []*FindResponseModel {
	var res = []*FindResponseModel{}

	for rows.Next() {
		var model = new(FindResponseModel)
		var err = rows.Scan(
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
