package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	//"google.golang.org/grpc/status"
	//"google.golang.org/grpc/codes"
	pb "github.com/wechselstrom/schafkopfturnier/proto"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedSchafkopfServer
}

func (s *server) SendMsg(ctx context.Context, in *pb.Message) (*pb.Empty, error) {
	log.Printf("Received: %v", in.GetText())
	return &pb.Empty{}, nil
}

func (s *server) KarteSpielen(ctx context.Context, in *pb.SpielKarteRequest) (*pb.Empty, error) {
	log.Printf("Value: %d", wert_zu_augen[in.GetKarte().GetWert()])
	spiel := pb.Spiel{Spieltyp:pb.Spieltyp_Solo, Farbe:pb.Farbe_Herz}
	log.Printf("Ist Trumpf: %v", istTrumpf(spiel, *in.GetKarte()))
	//err := status.Error(codes.NotFound, "id was not found")
	//return nil, err
	return &pb.Empty{}, nil
}


func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSchafkopfServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
