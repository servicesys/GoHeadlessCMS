package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"os"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "valter"
	password = "valter"
	dbname   = "app_sistema"
)

func Connect() *pgx.Conn {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := pgx.Connect(context.Background(), psqlInfo)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		panic(err)
	}
	fmt.Println("Successfully connected!")
	fmt.Println(db.Config())

	return db

}

func doExecute(db *pgx.Conn, query string, args ...interface{}) error {

	_, err := db.Exec(context.Background(), query, args...)

	return err
}
