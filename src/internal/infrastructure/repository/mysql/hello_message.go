package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

type repository struct {
	listName []string
	db       *sqlx.DB
}

func NewHelloMessageRepo() *repository {
	db, err := sqlx.Connect("mysql", "admin:admin@(localhost:3306)/messages")
	if err != nil {
		log.Fatalln(err)
	}
	repo := &repository{}
	repo.db = db
	return repo
}

func (r *repository) Save(name string) error {
	tx := r.db.MustBegin()
	tx.MustExec("INSERT INTO messages VALUES (default,'" + name + "')")
	err := tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetList() []string {
	rows, _ := r.db.Queryx("SELECT message FROM messages")
	for rows.Next() {
		err := rows.StructScan(&r.listName)
		if err != nil {
			log.Fatalln(err)
		}
	}
	return r.listName
}
