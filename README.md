# Go gRPC Report Generation Service with Scheduled Cron Job

## 🧾 Overview

This project is a small gRPC-based Go service designed to:
- Provide a `GenerateReport` gRPC endpoint
- Periodically generate reports for predefined users using a cron job
- Store reports in-memory
- Include a `HealthCheck` gRPC endpoint
- Log all operations with timestamps

It follows best practices for modular Go development and is fully testable via unit tests, a gRPC client, and CLI tools like Evans.

---

## ✅ How to Run the Service

### 1. Install Dependencies
Ensure Go is installed, then run:

```bash
go mod tidy
```

### 2. Generate gRPC Code
Install protoc: https://github.com/protocolbuffers/protobuf/releases

and run:
protoc --go_out=. --go-grpc_out=. *.proto

### 3. Start the Server

go run main.go

The gRPC server will start on localhost:8000 and begin executing a cron job every 10 seconds to generate reports for user1, user2, and user3.

## How to Test gRPC Calls

### A. Unit Tests

go test ./server

This runs:

 1) TestGenerateReport: validates correct report generation and memory storage

 2) TestHealthCheck: validates the health endpoint response

### B. gRPC Client Test 

go run client/main.go




