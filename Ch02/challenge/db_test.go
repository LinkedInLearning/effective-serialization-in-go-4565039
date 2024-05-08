package db

import (
	"path"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDB(t *testing.T) {
	dbPath := path.Join(t.TempDir(), "users.db")
	db, err := Open(dbPath)
	require.NoError(t, err, "open")
	defer db.Close()

	u := User{
		Login: "elliot",
		Roles: []Role{RoleReader, RoleWriter},
	}
	err = db.Put(u)
	require.NoError(t, err, "put")

	u2, err := db.Get(u.Login)
	require.NoError(t, err, "get")
	require.Equal(t, u, u2)
}
