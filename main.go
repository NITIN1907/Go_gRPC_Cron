package main

import (
	"grpc/cron"
	proto "grpc/proto"
	"grpc/server"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reportServer := &server.ReportServer{Reports: make(map[string]string)}
	proto.RegisterReportServiceServer(s, reportServer)

	go cron.StartCronJob(reportServer)

	log.Println("gRPC server started on port:  :8000")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}

}
