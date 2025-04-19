package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func GetConnection() *sql.DB {
	env := viper.New()
	env.SetConfigFile(".env")
	env.AddConfigPath("../")

	err := env.ReadInConfig()
	if err != nil {
		panic(err)
	}

	dbUsername := env.GetString("DB_USERNAME")
	dbPassword := env.GetString("DB_PASSWORD")
	dbPort := env.GetString("DB_PORT")
	dbName := env.GetString("DB_NAME")

	db, err := sql.Open("mysql", dbUsername+":"+dbPassword+"@tcp(localhost:"+dbPort+")/"+dbName)
	if err != nil {
		panic(err)
	}

	return db
}
