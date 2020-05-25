package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	elastic "gopkg.in/olivere/elastic.v7"
	pb "grpc_es_streaming/app/interface/rpc/protos"
	"log"
	"net"
)
const (
	grpcPort = ":50051"
)


type server struct{}
func (s server) GetAssetChanges(in *pb.AssetRequest, assetServer pb.Asset_GetAssetChangesServer) error {
	//ctx := context.Background()
	client, _ :=  elastic.NewClient(elastic.SetURL("http://localhost:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false))


	searchValue := in.Name
	fmt.Println("ES initialized...")
	log.Println("start new server")
	ctx := assetServer.Context()
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		searchSource := elastic.NewSearchSource()
		searchSource.Query(elastic.NewMatchQuery("Name", searchValue))
		searchService := client.Search().Index("asset").SearchSource(searchSource)
		searchResult, err := searchService.Do(ctx)
		if err != nil {
			fmt.Println("[ProductsES][GetPIds]Error=", err)
		}

		resp := pb.AssetResponse{NumberOfAssets: int32(searchResult.TotalHits())}
		if err := assetServer.Send(&resp); err != nil {
			log.Printf("send error %v", err)
		}
		log.Printf("send new assetResponse=%d", 1)
	}
}
func main() {
	listen, err := net.Listen("tcp", ":50005")
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
		return
	}
	grpcServer := grpc.NewServer()
	pb.RegisterAssetServer(grpcServer, server{})
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
