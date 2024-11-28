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

// Hash 3790
// Hash 6185
// Hash 6368
// Hash 7066
// Hash 9967
// Hash 5585
// Hash 3384
// Hash 6859
// Hash 8807
// Hash 7684
// Hash 6240
// Hash 2486
// Hash 4332
// Hash 8703
// Hash 1001
// Hash 5864
// Hash 6406
// Hash 8749
// Hash 9637
// Hash 2905
// Hash 4491
// Hash 6885
// Hash 2344
// Hash 1115
// Hash 6616
// Hash 5223
// Hash 9661
// Hash 9629
// Hash 7635
// Hash 3871
// Hash 3777
// Hash 4728
// Hash 8336
// Hash 3838
// Hash 9706
// Hash 2182
// Hash 3447
// Hash 6676
// Hash 9148
// Hash 5272
// Hash 9010
// Hash 8192
// Hash 5758
// Hash 6074