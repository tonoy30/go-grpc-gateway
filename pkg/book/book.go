package book

import (
	"context"
	"github.com/tonoy30/recap-grpc/pb"
)

type Server struct {
	pb.UnimplementedBookServiceServer
}

func (s *Server) GetBook(ctx context.Context, req *pb.GetBookRequest) (*pb.GetBookResponse, error) {
	return &pb.GetBookResponse{
		Id: req.GetId(),
		Title: "The Create method takes a parent resource name, a resource, and zero or more parameters." +
			" It creates a new resource under the specified parent, and returns the newly created resource.",
	}, nil
}
