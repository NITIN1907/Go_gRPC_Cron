package main

import (
	"context"
	"log"
	"time"

	proto "grpc/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewReportServiceClient(conn)

	// Test GenerateReport
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := client.GenerateReport(ctx, &proto.UserRequest{UserID: "client_test_user"})
	if err != nil {
		log.Fatalf("GenerateReport failed: %v", err)
	}
	log.Printf("Report generated: %s\n", resp.ReportID)

	// Test HealthCheck
	healthResp, err := client.HealthCheck(ctx, &proto.HealthRequest{})
	if err != nil {
		log.Fatalf("HealthCheck failed: %v", err)
	}
	log.Printf("HealthCheck status: %s\n", healthResp.Status)
}
