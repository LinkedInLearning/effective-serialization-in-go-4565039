package db

import (
	"bytes"
	"encoding/gob"
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
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(u); err != nil {
		return err
	}

	err := db.conn.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket(bucketName)
		return b.Put([]byte(u.Login), buf.Bytes())
	})
	if err != nil {
		return err
	}

	return nil
}

// Get gets a user from the database.
func (db *DB) Get(login string) (User, error) {
	var data []byte

	err := db.conn.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(bucketName)
		data = b.Get([]byte(login))
		return nil
	})

	if err != nil {
		return User{}, err
	}

	if data == nil {
		return User{}, fmt.Errorf("%q - not found", login)
	}

	var u User
	if err := gob.NewDecoder(bytes.NewReader(data)).Decode(&u); err != nil {
		return User{}, err
	}

	return u, nil
}
