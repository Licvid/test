package main

import (
	"log"
	"net"
	"os/exec"

	pb "server/proto"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"google.golang.org/grpc/reflection"
)

const (
	port = ":50055"
)

type server struct{}

func execute(cmds ...string) []byte {
	var log []byte
	for _, cmd := range cmds {
		out, err := exec.Command("bash", "-c", cmd).Output()
		if err != nil {
			panic("some error found")
		}
		log = append(log, out...)
	}
	return log
}

func (s *server) AppList(ctx context.Context, in *pb.AppListRequest) (*pb.AppListReply, error) {
	out := execute("sudo docker ps")
	return &pb.AppListReply{Result: out}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSdronServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}

}
