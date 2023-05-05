package main

import (
	"github.com/fahad-md-kamal/fiber-blogs/configs"
	"github.com/fahad-md-kamal/fiber-blogs/database"
	"github.com/fahad-md-kamal/fiber-blogs/migrations"
	"github.com/fahad-md-kamal/fiber-blogs/server"
)

func main() {
	if err := configs.LoadEnvs(); err != nil {
		panic(err.Error())
	}
	if err := database.DbConfig(); err != nil {
		panic(err.Error())
	}
	migrations.MigrateChanges()
	server.SetupAndListen()
}
