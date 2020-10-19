package config

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
)

var Config *Setter

type Setter struct {
  DB *sql.DB
	env map[string]string
}

func GetConfig() *Setter{
	env := LoadEnv()
  db := GetConnection()
  Config = &Setter{db, env}
  return Config
}

// Load the .env file in the current directory
func LoadEnv() map[string]string {
	val, err := godotenv.Read()
	if err != nil {
    log.Print("Error loading .env file")
  }

	return val
}

func (*Setter) Getenv(key string) string {
	return LoadEnv()[key]
}
