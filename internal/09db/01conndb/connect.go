package conndb

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func Connect() {
	if DB != nil {
		return
	}
	var err error
	DB, err = sql.Open("pgx", "host=localhost port=5432 user=testuser dbname=testdb password=pass sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func Ping() {
	Connect()
	if err := DB.Ping(); err != nil {
		log.Fatalf("ping database: %v", err)
	}
}

type User struct {
	UserID    string
	UserName  string
	CreatedAt sql.NullTime
}

func QueryUsers() []*User {
	Connect()
	rows, err := DB.QueryContext(context.TODO(), `SELECT user_id, user_name, created_at FROM users ORDER BY user_id;`)
	if err != nil {
		log.Fatalf("query all users: %v", err)
	}
	defer rows.Close()

	var users []*User
	for rows.Next() {
		var (
			userID, userName string
			createdAt        sql.NullTime
		)
		if err := rows.Scan(&userID, &userName, &createdAt); err != nil {
			log.Fatalf("scan the user: %v", userID)
		}
		users = append(users, &User{
			UserID:    userID,
			UserName:  userName,
			CreatedAt: createdAt,
		})
	}
	if err := rows.Close(); err != nil {
		log.Fatalf("rows close: %v", err)
	}
	if err := rows.Err(); err != nil {
		log.Fatalf("scan uses: %v", err)
	}
	return users
}

func Transaction() error {
	Connect()

	ctx := context.TODO()
	tx, err := DB.Begin()
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, "UPDATE users SET user_name = $1 WHERE user_id = $2", "New Name", "1")
	if err != nil {
		return fmt.Errorf("update user: %w", err)
	}

	return tx.Commit()
}

type txAdmin struct {
	db *sql.DB
}

type Service struct {
	txAdmin
}

func NewService(db *sql.DB) *Service {
	return &Service{
		txAdmin: txAdmin{
			db: db,
		},
	}
}

func (t *txAdmin) Transaction(ctx context.Context, f func(ctx context.Context, tx *sql.Tx) error) error {
	tx, err := t.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := f(ctx, tx); err != nil {
		return fmt.Errorf("transaction query failed: %w", err)
	}

	return tx.Commit()
}
