package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	dbURL := "abc"
	dbName := "abc"
	dbUser := "abc"
	dbPsswrd := "abc"

	connectStr := fmt.Sprintf(`postgres://%v:%v@%v/%v`, dbUser, dbPsswrd, dbURL, dbName)
	conn, err := sql.Open("postgres", connectStr)
	if err != nil {
		log.Println("DB Connect Init::Err::", err)
	}
	return conn
}
