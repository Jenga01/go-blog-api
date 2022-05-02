package main

import (
	"first/server"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "gorm.io/driver/postgres"
)

func main() {
	server.Run()
}
