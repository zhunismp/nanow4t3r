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

	err = db.AutoMigrate(&domain.BottledWater{})
	if err != nil {
		panic("failed to migrate database")
	}

	configureDBConnection(db)

	return db
}

func configureDBConnection(gorm *gorm.DB) {
	sqlDB, err := gorm.DB()
	if err != nil {
		fmt.Println("Error getting *sql.DB object:", err)
		return
	}

	// config connection pool
	sqlDB.SetMaxOpenConns(100)  // Maximum number of open connections
	sqlDB.SetMaxIdleConns(10)   // Maximum number of idle connections
	sqlDB.SetConnMaxLifetime(0) // Maximum amount of time a connection can be reused (0 means no limit)

}

func buildConnectionString(DBConfig config.DBConfig) string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		DBConfig.HOST, DBConfig.PORT, DBConfig.USERNAME, DBConfig.PASSWORD, DBConfig.NAME,
	)
}
