package utils

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Asteroid struct {
	Id      int    `json:"id"`
	Version string `json:"version"`
	SQL     string `json:"sql"`
}

// this will select requested update from database
func Select() (Asteroid, bool) {
	HOST := GetEnv("HOST")
	USERNAME := GetEnv("USERNAME")
	PASSWORD := GetEnv("PASSWORD")
	DATABASE := GetEnv("DATABASE")

	db, err := sql.Open("mysql", USERNAME+":"+PASSWORD+"@tcp("+HOST+")/"+DATABASE)

	if err != nil {
		panic(err.Error())
	}

	res, err := db.Query("SELECT `id`, `version`, `sql` FROM probat_updates WHERE request_update = 1 AND updated = 0 LIMIT 1")

	if err != nil {
		panic(err.Error())
	}

	var update Asteroid

	if res.Next() {
		err = res.Scan(&update.Id, &update.Version, &update.SQL)
		if err != nil {
			panic(err.Error())
		}

		return update, true
	}

	return update, false

}
// this will update requested update in database
func Update(id int) {
	HOST := GetEnv("HOST")
	USERNAME := GetEnv("USERNAME")
	PASSWORD := GetEnv("PASSWORD")
	DATABASE := GetEnv("DATABASE")

	db, err := sql.Open("mysql", USERNAME+":"+PASSWORD+"@tcp("+HOST+")/"+DATABASE)

	if err != nil {
		panic(err.Error())
	}

	_, err = db.Exec("UPDATE probat_updates SET updated = 1 WHERE id = ?", id)

	if err != nil {
		panic(err.Error())
	}
}
