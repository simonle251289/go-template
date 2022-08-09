package postgres

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"template/configs"
)

func PgConnection(config *configs.Config) *gorm.DB {
	//Load main database
	pg := config.Database.Postgres
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d", pg.Host, pg.User, pg.Password, pg.Database, pg.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	return db
}
