package g

import (
	"sync"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

var (
	dbr    *sqlx.DB
	dbw    *sqlx.DB
	dbOnce sync.Once
)

func DB(read ...bool) *sqlx.DB {
	dbOnce.Do(func() {
		var err error
		dbw, err = sqlx.Open("pgx", Cfg().Database.Write)
		if err != nil {
			panic(err)
		}
		if Cfg().Database.Read != "" && Cfg().Database.Read != Cfg().Database.Write {
			dbr, err = sqlx.Open("pgx", Cfg().Database.Read)
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

func RDB() *sqlx.DB {
	return DB(true)
}

func WDB() *sqlx.DB {
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
