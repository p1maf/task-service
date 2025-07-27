package main

import (
	"github.com/p1maf/task-service/internal/database"
	"github.com/p1maf/task-service/internal/task"
	transportgrpc "github.com/p1maf/task-service/internal/transport/grpc"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	database.InitDB()
	repo := task.NewRepository(database.DB)
	service := task.NewService(repo)

	userClient, conn, err := transportgrpc.NewUserClient("localhost:50051")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	go func() {
		log.Println("pprof on :6060")
		if err := http.ListenAndServe(":6060", nil); err != nil {
			log.Fatalf("pprof failed: %v", err)
		}
	}()
	if err := transportgrpc.RunGRPC(service, userClient); err != nil {
		log.Fatal(err)
	}
}
