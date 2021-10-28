package main

import (
	"api/internal/api/config"
)

func main() {
	c := config.LoadConfiguration()
	r := config.NewResolver(c)
	logger := r.ResolveLogger()

	db := r.ResolveSQLDatabase()
	db.Migrate()

	api := r.ResolveApiServer()
	if err := api.Run(); err != nil {
		logger.Panic(err)
	}
}
