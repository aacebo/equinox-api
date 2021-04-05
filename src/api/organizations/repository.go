package organizations

import (
	"database/sql"

	"github.com/aacebo/equinox-api/src/db"
	"github.com/aacebo/equinox-api/src/page"
)

type Repository struct {
	_db  *sql.DB
	_sql map[string]string
}

func NewRepository(conn *sql.DB) *Repository {
	var r = new(Repository)
	var q, err = db.LoadScripts("organizations")

	if err != nil {
		log.Error(err)
	}

	r._db = conn
	r._sql = q

	return r
}

func (r *Repository) Find(p *page.Page) ([]*Model, int) {
	var rows, rerr = r._db.Query(r._sql["find"], p.Like(), p.Skip(), p.PerPage)

	if rerr != nil {
		log.Error(rerr)
	}

	defer rows.Close()

	var count, cerr = r._db.Query(r._sql["find_count"], p.Like())

	if cerr != nil {
		log.Error(cerr)
	}

	defer count.Close()
	var total int

	count.Next()
	count.Scan(&total)

	return NewModels(rows), total
}

func (r *Repository) FindBySlug(slug string) *Model {
	var rows, err = r._db.Query(r._sql["find_by_slug"], slug)

	if err != nil {
		log.Error(err)
	}

	defer rows.Close()

	return NewModel(rows)
}

func (r *Repository) Mock() *Model {
	var model = Mock()
	var _, err = r._db.Exec(r._sql["create"], model.ID, model.Slug, model.Name, model.CreatedAt, model.UpdatedAt)

	if err != nil {
		log.Error(err)
	}

	return model
}
