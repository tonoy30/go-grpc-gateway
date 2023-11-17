package calculator

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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

func (c *Server) Sub(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	return &pb.Response{
		Result: req.GetNum1() - req.GetNum2(),
	}, nil
}

func (c *Server) Multiply(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	return &pb.Response{
		Result: req.GetNum1() * req.GetNum2(),
	}, nil
}

func (c *Server) Divide(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	if req.GetNum2() == 0 {
		return nil, status.Error(codes.InvalidArgument,
			fmt.Sprintf("%v can't be divied by %v", req.GetNum1(), req.GetNum2()),
		)
	}
	return &pb.Response{
		Result: req.GetNum1() / req.GetNum2(),
	}, nil
}
