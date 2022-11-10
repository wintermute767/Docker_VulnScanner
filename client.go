package main

import (
	api "VulnScanner/pkg/api/v1"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":10501", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("not conn %s", err)
	}
	defer conn.Close()
	c := api.NewNetVulnServiceClient(conn)
	msg := api.CheckVulnRequest{
		Targets: []string{"127.0.0.1"},
		TcpPort: []int32{22, 26, 80, 443, 587},
	}

	response, err := c.CheckVuln(context.Background(), &msg)
	if err != nil {
		log.Fatalf("error %s", err)
	}
	log.Println(response)
}
