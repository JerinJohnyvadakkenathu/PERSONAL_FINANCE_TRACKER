package main

import (
	models "FINANCE/Models"
	postgresql "FINANCE/Postgresql"
)

func main() {
	postgresql.ConnectDB()
	postgresql.DB.AutoMigrate(&models.User{})

}
