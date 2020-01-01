package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/salmander/go-grpc-tutorial/userservice"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50052"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("error connecting to %s, err: %v", address, err)
	}
	defer conn.Close()
	c := pb.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Usage:
	// go run userservice/main.go id=9
	// go run userservice/main.go uuid=batman-911
	// go run userservice/main.go email=imbatman@justiceleague.com
	// go run userservice/main.go nectar=911

	idPtr := flag.Int("id", 0, "SmartShop identity id")
	//uuidPtr := flag.String("uuid", "batman-911", "Sainsbury's identity UUID")
	flag.Parse()
	if idPtr != nil && *idPtr > 0 {
		log.Printf("Getting customer by SmartShop ID: %v", *idPtr)
		user, err := c.GetCustomerById(ctx, &pb.CustomerByIdRequest{Id: int32(*idPtr)})
		if err != nil {
			log.Printf("error returned from GetCustomerById: %v, Err: %v", *idPtr, err)
			return
		}
		log.Printf("User found by SmartShop ID: %v, %v", *idPtr, *user)
	}
}
