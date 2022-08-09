package appcontext

import (
	"gorm.io/gorm"
	"template/configs"
	"template/databases/postgres"
)

type AppContext interface {
	GetMainPostgresConnection() *gorm.DB
	GetConfig() *configs.Config
}

type appContext struct {
	pgDb   *gorm.DB
	config *configs.Config
}

func NewAppContext(config *configs.Config) *appContext {
	pgDb := postgres.PgConnection(config)
	return &appContext{pgDb: pgDb, config: config}
}

func (a *appContext) GetMainPostgresConnection() *gorm.DB {
	return a.pgDb
}

func (a *appContext) GetConfig() *configs.Config {
	return a.config
}
