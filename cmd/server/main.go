package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"log"
	"net"
	apisrv "testTask/internal/app/api"
	"testTask/internal/database"
	"testTask/internal/services/productservice"
	desc "testTask/pkg/api"
	"time"
)

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:example@127.0.0.1:27017/"))
	if err != nil {
		log.Fatalf("couldn't create mongo client: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := client.Connect(ctx); err != nil {
		log.Fatalf("can't connect to mongo: %v", err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatalf("can't disconnect mongo: %v", err)
		}
	}()

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("can't ping mongo: %v", err)
	}
	log.Print("mongo is started")

	db := database.NewStorage(client, "public")
	srv := apisrv.NewApi(productservice.NewProvider(db))
	s := grpc.NewServer()

	desc.RegisterTestTaskServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("can't listen: %v", err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}