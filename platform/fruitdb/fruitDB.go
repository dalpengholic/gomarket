package fruitdb

import "database/sql"

type FruitInfo struct {
	DB *sql.DB
}

func (f *FruitInfo) Add(item Item) {
	stmt, _ := f.DB.Prepare(`
		INSERT INTO fruitinfo (ID, Name, Price) values (?, ?, ?)
	`)
	stmt.Exec(item.ID, item.Name, item.Price)
}

func NewFruitInfo(db *sql.DB) *FruitInfo {
	stmt, _ := db.Prepare(`
	CREATE TABLE IF NOT EXISTS "fruitinfo" (
		"ID"	INTEGER NOT NULL UNIQUE,
		"Name"	TEXT,
		"Price" REAL,
		PRIMARY KEY("ID" AUTOINCREMENT)
	);
	`)
	stmt.Exec()

	return &FruitInfo{
		DB: db,
	}
}
