package postgres

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"template/configs"
	"time"
)

func PgConnection(config *configs.Config) *gorm.DB {
	var newLogger logger.Interface
	if config.Debug {
		newLogger = logger.Default.LogMode(logger.Info)
	}
	//Load main database
	pg := config.Database.Postgres
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d", pg.Host, pg.User, pg.Password, pg.Database, pg.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})
	if err != nil {
		log.Fatalln(err)
	}
	return db
}
