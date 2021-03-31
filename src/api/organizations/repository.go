package organizations

import (
	"database/sql"
	"log"

	"github.com/aacebo/equinox-api/src/db"
)

type Repository struct {
	_db  *sql.DB
	_sql map[string]string
}

func NewRepository(conn *sql.DB) *Repository {
	var r = new(Repository)
	var q, qe = db.LoadScripts("organizations")

	if qe != nil {
		log.Fatal(qe)
	}

	r._db = conn
	r._sql = q

	return r
}

func (r *Repository) Find() []*Model {
	var rows, err = r._db.Query(r._sql["find"])

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	return r._RowsToArray(rows)
}

func (r *Repository) _RowsToArray(rows *sql.Rows) []*Model {
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
			log.Fatal(err)
		}

		res = append(res, model)
	}

	return res
}
