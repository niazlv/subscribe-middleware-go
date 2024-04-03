package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Subscibe struct {
	Id         string
	Subscribe1 string
	Subscribe2 string
	Next       string
}

func InitDB() *sql.DB {
	db, err := sql.Open("sqlite3", "storage/subscribe.db")
	if err != nil {
		panic(err)
	}
	if db == nil {
		panic("db nil")
	}
	return db
}

func CreateTable(db *sql.DB) {
	// create table if not exists
	sql_table := `
	CREATE TABLE IF NOT EXISTS subscribes(
		Id TEXT NOT NULL PRIMARY KEY,
		Subscribe1 TEXT NOT NULL ,
		Subscribe2 TEXT,
		Next TEXT,
		InsertedDatetime DATETIME
	);
	`

	_, err := db.Exec(sql_table)
	if err != nil {
		panic(err)
	}
}
func StoreSubscribe(db *sql.DB, subscribe Subscibe) {
	sql_additem := `
	INSERT OR REPLACE INTO subscribes(
		Subscribe1,
		Subscribe2,
		Next,
		Id,
		InsertedDatetime
	) values (?,?,?,?,CURRENT_TIMESTAMP)
	`
	stmt, err := db.Prepare(sql_additem)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(subscribe.Subscribe1, subscribe.Subscribe2, subscribe.Next, subscribe.Id)
	if err2 != nil {
		panic(err2)
	}
}

func StoreSubscribes(db *sql.DB, subscribes []Subscibe) {
	sql_additem := `
	INSERT OR REPLACE INTO subscribes(
		Subscribe1,
		Subscribe2,
		Next,
		Id,
		InsertedDatetime
	) values (?,?,?,?,CURRENT_TILESTAMP)
	`
	stmt, err := db.Prepare(sql_additem)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	for _, item := range subscribes {
		_, err2 := stmt.Exec(item.Subscribe1, item.Subscribe2, item.Next, item.Id)
		if err2 != nil {
			panic(err2)
		}
	}
}

func ReadSubscribes(db *sql.DB) []Subscibe {
	sql_readall := `
	SELECT Id, Subscribe1, Subscribe2, Next FROM subscribes
	ORDER BY datetime(InsertedDatetime) DESC
	`

	rows, err := db.Query(sql_readall)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var result []Subscibe
	for rows.Next() {
		item := Subscibe{}
		err2 := rows.Scan(&item.Id, &item.Subscribe1, &item.Subscribe2, &item.Next)
		if err2 != nil {
			panic(err2)
		}
		result = append(result, item)
	}
	return result
}

func ReadSubscribeByID(db *sql.DB, id string) (Subscibe, error) {
	sql_readone := `
	SELECT Id, Subscribe1, Subscribe2, Next FROM subscribes
	WHERE Id = ?
	`

	row := db.QueryRow(sql_readone, id)
	item := Subscibe{}
	err := row.Scan(&item.Id, &item.Subscribe1, &item.Subscribe2, &item.Next)
	if err != nil {
		return Subscibe{}, err
	}

	return item, nil
}
