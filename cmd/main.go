package main

import (
	"context"
	"log"

	"lamoda_task/internal/config"
	"lamoda_task/internal/db"
	"lamoda_task/internal/server"
	"lamoda_task/internal/store"
	"lamoda_task/internal/web"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := config.LoadConfig(ctx)
	if err != nil {
		log.Fatalf("Cannot load config: %v", err)
	}
	log.Println("Config initialized!")

	dbInst, err := db.NewPgx(cfg.DB)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	log.Println(ctx, "Connected to database")

	defer func() {
		err := dbInst.Close()
		if err != nil {
			log.Fatalf("can not close database client: %v \n", err)
		}
	}()

	storeManagement := store.NewStoreManagement(dbInst)

	router := server.InitHttpRouter()
	web.InitRoutes(router, storeManagement)
	log.Println("HTTP routes initialized")

	server.RunHttpServer(ctx, router, cfg.HTTPServer)
}
