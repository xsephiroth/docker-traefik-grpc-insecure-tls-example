package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"time"
	pb "traefikgrpc/proto"
)

var (
	server  = flag.String("server", "server:8000", "server address")
	withTls = flag.Bool("tls", false, "with certs")
)

func main() {
	flag.Parse()

	var opts []grpc.DialOption
	if !*withTls {
		opts = []grpc.DialOption{grpc.WithInsecure()}
		log.Println("with insecure")
	} else {
		opts = withTlsOptions()
		log.Println("with tls")
	}

	conn, err := grpc.Dial(*server, opts...)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = conn.Close()
	}()

	cli := pb.NewTraefikGRPCProxyClient(conn)

	req := time.Now().String()
	log.Println("request to server: ", req)

	resp, err := cli.ProxyMe(context.Background(), &pb.ProxyMeRequest{
		Req: req,
	})
	if err != nil {
		panic(err)
	}

	log.Println("received from server: ", resp.GetResp())
}

func withTlsOptions() []grpc.DialOption {
	creds, err := credentials.NewClientTLSFromFile(
		"/certs/srv.pem",
		"test.com", // # Common Name (e.g. server FQDN or YOUR name) []: test.com
	)
	if err != nil {
		panic(err)
	}

	return []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
	}
}
