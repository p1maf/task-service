package main

import (
	"fmt"
	"github.com/p1maf/task-service/internal/database"
	"github.com/p1maf/task-service/internal/task"
	transportgrpc "github.com/p1maf/task-service/internal/transport/grpc"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	fmt.Println("Starting server...")
	if err := transportgrpc.RunGRPC(service, userClient); err != nil {
		log.Fatal(err)
	}
}
