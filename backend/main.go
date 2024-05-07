package main

import (
	database "mmr/backend/db"
	server "mmr/backend/server"
)

func main() {
	database.Init()
	server.Init()
}
