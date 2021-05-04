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
	var rows, rerr = self._db.Query(self._sql["find"], p.Like(), p.Skip(), p.PerPage)

	if rerr != nil {
		log.Error.Fatal(rerr)
	}

	defer rows.Close()

	var count, cerr = self._db.Query(self._sql["find_count"], p.Like())

	if cerr != nil {
		log.Error.Fatal(cerr)
	}

	defer count.Close()
	var total int

	count.Next()
	count.Scan(&total)

	return NewModels(rows), total
}

func (self *Repository) FindById(id string) *Model {
	var rows, err = self._db.Query(self._sql["find_by_id"], id)

	if err != nil {
		log.Error.Fatal(err)
	}

	defer rows.Close()

	return NewModel(rows)
}

func (self *Repository) FindBySlug(slug string) *Model {
	var rows, err = self._db.Query(self._sql["find_by_slug"], slug)

	if err != nil {
		log.Error.Fatal(err)
	}

	defer rows.Close()

	return NewModel(rows)
}

func (self *Repository) Upsert(id string, slug string, name string, createdAt time.Time, updatedAt time.Time) {
	var existing = self.FindById(id)
	var cmd string = "create"

	if existing != nil {
		cmd = "update"
	}

	var _, err = self._db.Exec(
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
