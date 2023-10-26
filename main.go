package main

import (
	"assignment1/database"
	"assignment1/router"
)

var (
	PORT = ":8080"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(PORT)
}
