package calculator

import (
	"context"

	"github.com/tonoy30/recap-grpc/pb"
)

type Server struct {
	pb.UnimplementedCalculatorServiceServer
}

func (c *Server) Add(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	return &pb.Response{
		Result: req.GetNum1() + req.GetNum2(),
	}, nil
}
