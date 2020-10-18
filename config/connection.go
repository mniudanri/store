package config

import (
  "fmt"
	"database/sql"
	_ "github.com/lib/pq"
)

func GetConnection() *sql.DB{
  host     := Config.Getenv("DB_HOST")
  port     := Config.Getenv("DB_PORT")
  user     := Config.Getenv("DB_USER")
  password := Config.Getenv("DB_PASSWORD")
  dbname   := Config.Getenv("DB_NAME")

  psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)

  // NOTE: connection to postgres with local development
  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }

  return db
}
