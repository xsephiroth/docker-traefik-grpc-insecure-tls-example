package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "traefikgrpc/proto"
)

func main() {
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	pb.RegisterTraefikGRPCProxyServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}

type server struct {}

func (s *server) ProxyMe(ctx context.Context, req *pb.ProxyMeRequest) (*pb.ProxyMeResponse, error) {
	log.Println("received from client: ", req.GetReq())

	resp := "RESP "+ req.GetReq()
	log.Println("response to client: "+ resp)

	return &pb.ProxyMeResponse{
		Resp: resp,
	}, nil
}
