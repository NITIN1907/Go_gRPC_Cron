package server

import (
	"context"
	"testing"

	proto "grpc/proto"
)

func TestGenerateReport(t *testing.T) {
	s := &ReportServer{Reports: make(map[string]string)}

	req := &proto.UserRequest{UserID: "test_user"}
	resp, err := s.GenerateReport(context.Background(), req)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.ReportID == "" {
		t.Error("Expected a non-empty ReportID")
	}

	if val, ok := s.Reports["test_user"]; !ok || val != resp.ReportID {
		t.Errorf("Report not stored correctly in memory. Got: %v, Expected: %v", val, resp.ReportID)
	}
}

func TestHealthCheck(t *testing.T) {
	s := &ReportServer{Reports: make(map[string]string)}

	resp, err := s.HealthCheck(context.Background(), &proto.HealthRequest{})
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.Status != "OK" {
		t.Errorf("Expected status 'OK', got %v", resp.Status)
	}
}
