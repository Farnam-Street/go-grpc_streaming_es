package main
import (
	"fmt"
"net"
pb "grpc_es_streaming/protos"
"golang.org/x/net/context"
"google.golang.org/grpc"
"google.golang.org/grpc/reflection"
)
const (
	grpcPort = ":50051"
)
type Server struct{}
func (s *Server) GetAssetChanges(context context.Context, in *pb.AssetRequest) (*pb.AssetResponse, error) {
	return &pb.AssetResponse{NumberOfAssets: int32(len(in.Name))}, nil
}
func main() {
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
		return
	}
	grpcServer := grpc.NewServer()
	pb.RegisterAssetServer(grpcServer, &Server{})
	reflection.Register(grpcServer)
	grpcServer.Serve(listen)
}
