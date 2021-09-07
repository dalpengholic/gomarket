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

func (f *FruitInfo) Get() []Item {
	items := []Item{}
	rows, _ := f.DB.Query(`
	SELECT * FROM fruitinfo
	`)
	var id int
	var name string
	var price float64
	for rows.Next() {
		rows.Scan(&id, &name, &price)
		item := Item{
			ID:    id,
			Name:  name,
			Price: price,
		}
		items = append(items, item)
	}
	return items
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
