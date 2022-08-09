package postgres

import (
	"gorm.io/gorm"
	"log"
	"template/databases"
)

func ExampleAutoMigration(db *gorm.DB) {
	if err := db.AutoMigrate(&databases.TemplateModel{}).Error; err != nil {
		log.Fatalln(err())
	}
}
