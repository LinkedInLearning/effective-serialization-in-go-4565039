package db

import (
	"fmt"

	"go.etcd.io/bbolt"
)

type Role string

var (
	RoleReader Role = "reader"
	RoleWriter Role = "writer"
	RoleAdmin  Role = "admin"
)

type User struct {
	Login string
	Roles []Role
}

type DB struct {
	conn *bbolt.DB
}

var (
	bucketName = []byte("users")
)

// Open returns a new database.
func Open(dbPath string) (*DB, error) {
	conn, err := bbolt.Open(dbPath, 0666, nil)
	if err != nil {
		return nil, err
	}

	err = conn.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(bucketName)
		return err
	})
	if err != nil {
		return nil, err
	}

	return &DB{conn}, nil
}

// Close closes the database.
func (db *DB) Close() error {
	return db.conn.Close()
}

// Put puts a new user in the database.
func (db *DB) Put(u User) error {
	return fmt.Errorf("not implemented") // TODO
}

// Get gets a user from the database.
func (db *DB) Get(login string) (User, error) {
	return User{}, fmt.Errorf("not implemented") // TODO
}
