package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/tonoy30/recap-grpc/pkg/book"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"

	"github.com/tonoy30/recap-grpc/pb"
	"github.com/tonoy30/recap-grpc/pkg/calculator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatalln("failed to create listener: ", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterCalculatorServiceServer(s, &calculator.Server{})
	pb.RegisterBookServiceServer(s, &book.Server{})

	go func() {
		log.Printf("starting grpc server at port: %v\n", 8080)
		if err := s.Serve(lis); err != nil {
			log.Fatalf("error serving grpc: %v\n", err)
		}
	}()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err = pb.RegisterCalculatorServiceHandlerFromEndpoint(ctx, mux, "localhost:8080", opts)
	err = pb.RegisterBookServiceHandlerFromEndpoint(ctx, mux, "localhost:8080", opts)

	go func() {
		log.Println("serving grpc-gateway on http://localhost:8081")
		if err := http.ListenAndServe(":8081", mux); err != nil {
			log.Fatalf("error serving grpc gateway: %v\n", err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

	s.Stop()
	_ = lis.Close()
}
