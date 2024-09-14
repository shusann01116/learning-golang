package conndb_test

import (
	"testing"

	_ "github.com/jackc/pgx/v5"
	conndb "github.com/shusann01116/learning-golang/internal/09db/01conndb"
)

func TestConnDB(t *testing.T) {
	conndb.Connect()
}

func TestPing(t *testing.T) {
	conndb.Ping()
}

func TestQuery(t *testing.T) {
	users := conndb.QueryUsers()
	for _, u := range users {
		t.Logf("%+v", u)
	}
}
