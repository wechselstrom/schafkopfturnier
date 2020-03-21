package main

import (
	"context"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	pb "github.com/wechselstrom/schafkopfturnier/proto"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewSchafkopfClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = c.SendMsg(ctx, &pb.Message{Text: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	_, err = c.SendMsg(ctx, &pb.Message{Text: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	karte := pb.Karte{Wert:pb.Wert_Neun, Farbe: pb.Farbe_Herz}
	_, err = c.KarteSpielen(ctx, &pb.SpielKarteRequest{Karte: &karte})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("sent all commands")
}
