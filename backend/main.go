package main

import (
	_ "ariga.io/atlas-provider-gorm/gormschema"
	"mmr/backend/config"
)

import (
	database "mmr/backend/db"
	server "mmr/backend/server"
)

func main() {
	config.LoadEnv()
	database.Init()
	server.Init()
}
