package db

import (
	"fmt"

	"github.com/zhunismp/nanow4t3r/services/product/core/domain"
	"github.com/zhunismp/nanow4t3r/services/product/infrastructures/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetNewGormDBInstance(DBConfig config.DBConfig) *gorm.DB {
	db, err := gorm.Open(postgres.Open(buildConnectionString(DBConfig)), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&domain.Product{})
	if err != nil {
		panic("failed to migrate database")
	}

	return db
}

func buildConnectionString(DBConfig config.DBConfig) string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		DBConfig.HOST, DBConfig.PORT, DBConfig.USERNAME, DBConfig.PASSWORD, DBConfig.NAME,
	)
}
