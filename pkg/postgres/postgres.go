package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
	"sync"
)

const driverName = "postgres"

// global db instance
var db *sql.DB

var (
	lock = &sync.Mutex{}
)

func GetDB(DbDsn string) (*sql.DB, error) {
	var err error
	if db == nil {
		lock.Lock()
		defer lock.Unlock()
		if db == nil {
			db, err = Connect(DbDsn)
			if err != nil {
				return nil, err
			}
		}
	}
	return db, nil
}

func Connect(DbDsn string) (*sql.DB, error) {
	db, err := sql.Open(driverName, DbDsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
