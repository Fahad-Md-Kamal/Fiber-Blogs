package database

import (
	"fmt"

	"github.com/fahad-md-kamal/fiber-blogs/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DbConfig() error {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Dhaka",
		configs.ENVs.DbHost, configs.ENVs.DbUser, configs.ENVs.DbPassword, configs.ENVs.DbName, configs.ENVs.DbPort)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}
