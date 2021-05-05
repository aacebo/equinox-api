package organizations

import (
	"database/sql"
	"time"

	"github.com/aacebo/equinox-api/src/db"
	"github.com/aacebo/equinox-api/src/log"
	"github.com/aacebo/equinox-api/src/page"
)

type Repository struct {
	_db  *sql.DB
	_sql map[string]string
}

func NewRepository(conn *sql.DB) *Repository {
	var self = new(Repository)
	var q, err = db.LoadScripts("organizations")

	if err != nil {
		log.Error.Fatal(err)
	}

	self._db = conn
	self._sql = q

	return self
}

func (self *Repository) Find(p *page.Page) ([]*Model, int) {
	var rows, err = self._db.Query(self._sql["find"], p.Like(), p.Skip(), p.PerPage)
	defer rows.Close()

	if err != nil {
		log.Error.Fatal(err)
	}

	count, err := self._db.Query(self._sql["find_count"], p.Like())
	defer count.Close()

	if err != nil {
		log.Error.Fatal(err)
	}

	var total int

	count.Next()
	count.Scan(&total)

	return NewModels(rows), total
}

func (self *Repository) FindById(id string) *Model {
	var rows, err = self._db.Query(self._sql["find_by_id"], id)
	defer rows.Close()

	if err != nil {
		log.Error.Fatal(err)
	}

	return NewModel(rows)
}

func (self *Repository) FindBySlug(slug string) *Model {
	var rows, err = self._db.Query(self._sql["find_by_slug"], slug)
	defer rows.Close()

	if err != nil {
		log.Error.Fatal(err)
	}

	return NewModel(rows)
}

func (self *Repository) Upsert(id string, slug string, name string, createdAt time.Time, updatedAt time.Time) {
	var existing = self.FindById(id)
	var cmd = "create"

	if existing != nil {
		cmd = "update"
	}

	_, err := self._db.Exec(
		self._sql[cmd],
		id,
		slug,
		name,
		createdAt,
		updatedAt,
	)

	if err != nil {
		log.Error.Fatal(err)
	}
}
