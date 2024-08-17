package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/bap2pecs/reth-exex-grpc-example/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedExExHandlerServer
}

func (s *server) HandleChainCommitted(ctx context.Context, req *pb.ChainCommittedRequest) (*pb.HandlerResponse, error) {
	log.Printf("Chain committed: start block %d, end block %d", req.StartBlock, req.EndBlock)
	return &pb.HandlerResponse{SendFinishedHeightEvent: true, FinishedHeight: req.EndBlock}, nil
}

func (s *server) HandleChainReorged(ctx context.Context, req *pb.ChainReorgedRequest) (*pb.HandlerResponse, error) {
	log.Printf("Chain reorged: old start %d, old end %d, new start %d, new end %d",
		req.OldStartBlock, req.OldEndBlock, req.NewStartBlock, req.NewEndBlock)
	return &pb.HandlerResponse{SendFinishedHeightEvent: true, FinishedHeight: req.NewEndBlock}, nil
}

func (s *server) HandleChainReverted(ctx context.Context, req *pb.ChainRevertedRequest) (*pb.HandlerResponse, error) {
	log.Printf("Chain reverted: start block %d, end block %d", req.StartBlock, req.EndBlock)
	return &pb.HandlerResponse{SendFinishedHeightEvent: true, FinishedHeight: req.StartBlock - 1}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterExExHandlerServer(s, &server{})
	fmt.Println("gRPC server listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
