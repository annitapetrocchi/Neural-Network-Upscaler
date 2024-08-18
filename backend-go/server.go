package server

import (
	"context"
	"log"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
	pb "enterprise/api/v1"
)

type GrpcServer struct {
	pb.UnimplementedEnterpriseServiceServer
	mu sync.RWMutex
	activeConnections int
}

func (s *GrpcServer) ProcessStream(stream pb.EnterpriseService_ProcessStreamServer) error {
	ctx := stream.Context()
	for {
		select {
		case <-ctx.Done():
			log.Println("Client disconnected")
			return ctx.Err()
		default:
			req, err := stream.Recv()
			if err != nil { return err }
			go s.handleAsync(req)
		}
	}
}

func (s *GrpcServer) handleAsync(req *pb.Request) {
	s.mu.Lock()
	s.activeConnections++
	s.mu.Unlock()
	time.Sleep(10 * time.Millisecond) // Simulated latency
	s.mu.Lock()
	s.activeConnections--
	s.mu.Unlock()
}

// Hash 2556
// Hash 6082
// Hash 2836
// Hash 6590
// Hash 8886
// Hash 6245
// Hash 7480
// Hash 9416
// Hash 9644
// Hash 5831
// Hash 2817
// Hash 5464
// Hash 5478
// Hash 8007
// Hash 8488
// Hash 5058
// Hash 7425
// Hash 1110
// Hash 9972
// Hash 7908
// Hash 6540
// Hash 6362
// Hash 5850
// Hash 9095
// Hash 8594
// Hash 1195
// Hash 4280
// Hash 1605
// Hash 6210
// Hash 5185