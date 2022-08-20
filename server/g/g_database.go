package g

import (
	"database/sql"
	"sync"

	_ "github.com/jackc/pgx/v5"
)

var (
	dbr    *sql.DB
	dbw    *sql.DB
	dbOnce sync.Once
)

func DB(read ...bool) *sql.DB {
	dbOnce.Do(func() {
		var err error
		dbw, err = sql.Open("pgx", Cfg().Database.Write)
		if err != nil {
			panic(err)
		}
		if Cfg().Database.Read != "" {
			dbr, err = sql.Open("pgx", Cfg().Database.Read)
			if err != nil {
				panic(err)
			}
		}
	})
	if len(read) > 0 && read[0] && dbr != nil {
		return dbr
	}
	return dbw
}

func RDB() *sql.DB {
	return DB(true)
}

func WDB() *sql.DB {
	return DB()
}

func DbClose() error {
	if err := dbw.Close(); err != nil {
		return err
	}
	if dbr != nil {
		return dbr.Close()
	}
	return nil
}
