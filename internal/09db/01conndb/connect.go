package conndb

import (
	"context"
	"database/sql"
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
