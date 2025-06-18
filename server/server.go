package server

import (
	"context"
	"fmt"
	proto "grpc/proto"
	"log"
	"sync"
	"time"
)

var mu sync.Mutex

type ReportServer struct {
	proto.UnimplementedReportServiceServer
	Reports map[string]string
}

func (s *ReportServer) GenerateReport(c context.Context, req *proto.UserRequest) (*proto.ReportResponse, error) {
	mu.Lock()
	defer mu.Unlock()

	reportID := fmt.Sprintf("report_%s_%d", req.UserID, time.Now().Unix())
	s.Reports[req.UserID] = reportID
	log.Printf("Generated report for user %s: %s", req.UserID, reportID)

	return &proto.ReportResponse{ReportID: reportID}, nil
}

func (s *ReportServer) HealthCheck(c context.Context, _ *proto.HealthRequest) (*proto.HealthResponse, error) {
	log.Println("Health check called")
	return &proto.HealthResponse{Status: "OK"}, nil
}
