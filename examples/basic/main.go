package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/philodi-dev/rdcpass-notification-service/golang-sdk/smsc"
)

func main() {
	settings := LoadSettings()
	if settings.AppID == "" || settings.SecretKey == "" {
		log.Fatal("SMSC_APP_ID and SMSC_SECRET_KEY are required")
	}

	client, err := settings.NewClient()
	if err != nil {
		log.Fatalf("create client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	runHealth(ctx, client)
	runQuickOTP(ctx, client, settings.Phone)
	runMultiStageSMS(ctx, client, settings.Phone)

	if settings.OTPCode != "" {
		runQuickVerify(ctx, client, settings.Phone, settings.OTPCode)
	}
}

func runHealth(ctx context.Context, client *smsc.Client) {
	health, err := client.Platform().Health(ctx)
	if err != nil {
		log.Fatalf("health: %v", err)
	}
	fmt.Printf("health: %s — %s\n", health.Status, health.Description)
}

func runQuickOTP(ctx context.Context, client *smsc.Client, phone string) {
	resp, err := client.Quick().SendOTP(ctx, phone)
	if err != nil {
		log.Fatalf("quick send otp: %v", err)
	}
	fmt.Printf("quick otp: request_id=%s status=%s\n", resp.RequestID, resp.Status)
}

func runMultiStageSMS(ctx context.Context, client *smsc.Client, phone string) {
	session, err := client.Auth().CreateSession(ctx)
	if err != nil {
		log.Fatalf("create session: %v", err)
	}
	fmt.Printf("session: expires_in=%ds\n", session.ExpiresIn)

	resp, err := client.SMS().Send(ctx, smsc.SendSMSRequest{
		Phone:   phone,
		Content: "Hello from RDCPASS Go SDK",
	})
	if err != nil {
		log.Fatalf("send sms: %v", err)
	}
	fmt.Printf("sms: request_id=%s status=%s\n", resp.RequestID, resp.Status)
}

func runQuickVerify(ctx context.Context, client *smsc.Client, phone, code string) {
	resp, err := client.Quick().VerifyOTP(ctx, phone, code)
	if err != nil {
		log.Fatalf("quick verify otp: %v", err)
	}
	fmt.Printf("quick verify: request_id=%s status=%s signature=%s\n",
		resp.RequestID, resp.Status, resp.Signature)
}
