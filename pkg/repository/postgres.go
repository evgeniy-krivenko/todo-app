package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	usersTable          = "users"
	todoListsTable      = "todo_lists"
	usersTodoListsTable = "users_todo_lists"
	todoItemsTable      = "todo_items"
	listsItemsTable     = "lists_items"
)

type Config struct {
	Host     string
	Post     string
	Username string
	DBname   string
	Password string
	SSLMode  string
}

func NewPostgresDB(cnf Config) (*sqlx.DB, error) {
	db, error := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cnf.Host, cnf.Post, cnf.Username, cnf.DBname, cnf.Password, cnf.SSLMode,
	))
	if error != nil {
		return nil, error
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
