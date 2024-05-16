package main

import (
	"fmt"
	"io"
	"mmr/backend/db/models"
	"os"

	_ "ariga.io/atlas-go-sdk/recordriver"
	"ariga.io/atlas-provider-gorm/gormschema"
)

func main() {
	stmts, err := gormschema.New("postgres").Load(&models.User{}, &models.Team{}, &models.Match{})
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to load gorm schema: %v\n", err)
		os.Exit(1)
	}
	_, _ = io.WriteString(os.Stdout, stmts)
}
