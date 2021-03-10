package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	dbURL := "ec2-18-211-97-89.compute-1.amazonaws.com:5432"
	dbName := "de3dgdv7gidb12"
	dbUser := "krtelnhvanlslj"
	dbPsswrd := "f300b8d256d3bb11c895c01d9e46d2881d2c63a5f0ebf2813d7672ad4ea46f7c"

	connectStr := fmt.Sprintf(`postgres://%v:%v@%v/%v`, dbUser, dbPsswrd, dbURL, dbName)
	conn, err := sql.Open("postgres", connectStr)
	if err != nil {
		log.Println("DB Connect Init::Err::", err)
	}
	return conn
}
