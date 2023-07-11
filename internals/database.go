package internals

import "fmt"

type DatabaseInterface interface {
	New(string) error

	Store() error
}

type Database struct {
	path string
}

func (db *Database) New(path string) error {
	db.path = path
	fmt.Println("New database created at ", db.path)
	return nil
}

func (db *Database) Store() error {
	fmt.Println("Database stored ")
	return nil
}
