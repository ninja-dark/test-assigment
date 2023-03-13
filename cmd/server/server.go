package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/ninja-dark/test-assigment/grpcService"
	playcase "github.com/ninja-dark/test-assigment/internal/handler/server/playCase"
	"github.com/ninja-dark/test-assigment/internal/infrastructure/api/db"
	"github.com/ninja-dark/test-assigment/internal/playlist"
	"github.com/pressly/goose"
	"google.golang.org/grpc"
	_ "github.com/lib/pq"
)

var dbConfig = &db.Config{
	Host:     "localhost",
	Port:     "5432",
	User:     "postgres",
	Password: "postgres",
	DBName: "",
}

func main() {

	//database
	poolConf, err := db.NewPoolConfig(dbConfig)
	if err != nil {
		log.Fatal("configuration error: %v", err)
	}
	pool, err := db.NewPool(poolConf)
	if err != nil {
		log.Fatal("failed to connect db: %v", err)
	}

	repo := db.NewRepo(pool)

	mig, err := sql.Open("postgres", poolConf.ConnString())
	if err != nil {
		log.Fatal("failed to open: %w", err)
	}
	err = mig.Ping()
	if err != nil {
		log.Fatal("failed goose: %v", err)
	}

	err = goose.Up(mig, "./migrations")
	if err != nil {
		log.Fatal("failed migration: %v", err)
	}
	//player
	playlist := playlist.NewPlaylist()

	//start grpc
	server := grpc.NewServer()
	plServer := playcase.NewServic(playlist, repo)
	grpcService.RegisterMusicPlaylistServer(server, plServer)

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
