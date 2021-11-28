package main

import (
	"database/sql"
	"log"
	models "simple_blog/models/sqlc"
	"simple_blog/pkg/setting"
	"simple_blog/routers"
)

func main() {
	config, err := setting.LoadSetting("./conf/")
	if err != nil {
		log.Fatalf("cannot load config, err: %v", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalf("cannot connect postgres, err: %v", err)
	}

	queries := models.New(conn)

	server, err := routers.NewServer(config, queries)
	if err != nil {
		log.Fatalf("cannot create server, err: %v", err)
	}
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}


