package cron

import (
	"context"
	proto "grpc/proto"
	"log"

	"github.com/robfig/cron/v3"
)

func StartCronJob(server proto.ReportServiceServer) {
	c := cron.New()
	users := []string{"user1", "user2", "user3"}

	c.AddFunc("@every 10s", func() {
		log.Println("[INFO] Running scheduled report generation...")
		GenerateReportsForAllUsers(server, users)
	})

	c.Start()
	log.Println("[INFO] Cron job started to run every 10 seconds.")
}

func GenerateReportsForAllUsers(server proto.ReportServiceServer, users []string) {
	for _, user := range users {
		res, err := server.GenerateReport(context.Background(), &proto.UserRequest{UserID: user})
		if err != nil {
			log.Printf("Error generating report for %s:%v\n", user, err)
			continue
		}
		log.Printf("Cron generated report for %s: %s\n", user, res.ReportID)
	}
}
