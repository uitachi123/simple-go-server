package db

import (
	memdb "github.com/hashicorp/go-memdb"
)

type User struct {
	Email string
	Name  string
}

type DBHandler struct {
	db *memdb.MemDB
}

func Init() (*DBHandler, error) {
	// Create the DB schema
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"user": &memdb.TableSchema{
				Name: "user",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Email"},
					},
					"name": &memdb.IndexSchema{
						Name:    "name",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Name"},
					},
				},
			},
		},
	}

	db, err := memdb.NewMemDB(schema)
	if err != nil {
		return nil, err
	}

	// Insert data
	txn := db.Txn(true)
	users := []*User{
		&User{"bob@test.com", "Bob"},
		&User{"alice@test.com", "Alice"},
	}
	for _, u := range users {
		if err := txn.Insert("user", u); err != nil {
			return &DBHandler{db}, err
		}
	}
	txn.Commit()
	return &DBHandler{db}, nil
}
