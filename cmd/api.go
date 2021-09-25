package main

import "api/internal/api/config"

func main() {
	c := config.LoadConfiguration()
	r := config.NewResolver(c)
	logger := r.ResolveLogger()
	api := r.ResolveApiServer()

	err := api.Run()

	if err != nil {
		logger.Panic(err)
	}
}
