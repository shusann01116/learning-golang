package conndb_test

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/jackc/pgx/v5"
	conndb "github.com/shusann01116/learning-golang/internal/09db/01conndb"
	"github.com/stretchr/testify/assert"
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

func TestTransaction(t *testing.T) {
	err := conndb.Transaction()
	if err != nil {
		t.Error(err)
	}
}

func TestTxAdminTransaction(t *testing.T) {
	db, err := sql.Open("pgx", "host=localhost port=5432 user=testuser dbname=testdb password=pass sslmode=disable")
	if err != nil {
		t.Fatalf("open psql connection: %v", err)
	}

	service := conndb.NewService(db)
	if err := service.Transaction(context.Background(), func(ctx context.Context, tx *sql.Tx) error {
		_, err := tx.Exec("INSERT INTO users (user_id, user_name) VALUES ($1, $2)", "4", "People")
		if err != nil {
			return fmt.Errorf("insert new user: %w", err)
		}
		return nil
	}); err != nil {
		t.Errorf("failed to insert new user: %v", err)
	}

	row := db.QueryRow("SELECT user_id, user_name FROM users WHERE user_id = $1", "4")
	var user conndb.User
	if err := row.Scan(&user.UserID, &user.UserName); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "4", user.UserID)
}
