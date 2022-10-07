package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type Manager struct {
	connection *sql.DB
}

var manager *Manager

func Singleton() *Manager {
	if manager == nil {
		manager = &Manager{
			connection: initiate(),
		}
	}
	return manager
}

func (manager *Manager) CloseConnection() {
	err := manager.connection.Close()
	if err != nil {
		return
	}
}

func (manager *Manager) GetConnection() *sql.DB {
	return manager.connection
}

func initiate() *sql.DB {
	db, err := sql.Open("sqlite3", "bookstore.sqlite")

	if err != nil {
		log.Fatal(err)
		return nil
	}

	return db
}
