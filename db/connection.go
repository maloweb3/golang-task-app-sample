package db

import (
	"app/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/kelseyhightower/envconfig"
)

func Connect() (*sql.DB, error) {
	cfg := &config.DbConfig{}
	err := envconfig.Process("", cfg)
	if err != nil {
		log.Println(err)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo", cfg.Host, cfg.User, cfg.Password, cfg.Name, cfg.Port)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		// エラー処理
		log.Println("not success", err)
	}

	return db, nil
}
