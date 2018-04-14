package room

import (
	"database/sql"
	"fmt"

	"github.com/davecgh/go-spew/spew"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type Db struct {
	Db  *sql.DB
	cfg Config
}

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

// New returns a Roach with the sql.DB set with the postgres
// DB connection string in the configuration
func New(cfg Config) (roach Db, err error) {
	if cfg.Host == "" || cfg.Port == "" || cfg.User == "" ||
		cfg.Password == "" || cfg.Database == "" {
		err = errors.Errorf(
			"All fields must be set (%s)",
			spew.Sdump(cfg))
		return
	}

	roach.cfg = cfg
	// The first argument corresponds to the driver name that the driver
	// (in this case, `lib/pq`) used to register itself in `database/sql`.
	// The next argument specifies the parameters to be used in the connection.
	// Details about this string can be seen at https://godoc.org/github.com/lib/pq
	db, err := sql.Open("postgres", fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		cfg.User, cfg.Password, cfg.Database, cfg.Host, cfg.Port))
	if err != nil {
		err = errors.Wrapf(err,
			"Couldn't open connection to postgre database (%s)",
			spew.Sdump(cfg))
		return
	}

	// Ping verifies if the connection to the database is alive or if a
	// new connection can be made.
	if err = db.Ping(); err != nil {
		err = errors.Wrapf(err,
			"Couldn't ping postgre database (%s)",
			spew.Sdump(cfg))
		return
	}

	roach.Db = db
	return
}

// Close performs the release of any resources that
// `sql/database` DB pool created. This is usually meant
// to be used in the exitting of a program or `panic`ing.
func (r *Db) Close() (err error) {
	if r.Db == nil {
		return
	}

	if err = r.Db.Close(); err != nil {
		err = errors.Wrapf(err,
			"Errored closing database connection",
			spew.Sdump(r.cfg))
	}

	return
}
