package organizations

import (
	"database/sql"

	log "github.com/sirupsen/logrus"
)

type Repository struct {
	_db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	var r = new(Repository)

	r._db = db

	return r
}

func (r *Repository) Find() []*Model {
	var rows, err = r._db.Query(`
	SELECT
		id,
		slug,
		name,
		created_at,
		updated_at
	FROM organizations
	`)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	return r._RowsToArray(rows)
}

func (r *Repository) _RowsToArray(rows *sql.Rows) (res []*Model) {
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
			log.Fatal(err)
		}

		res = append(res, model)
	}

	return res
}
