package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type EmailService struct{}

func NewEmailService() *EmailService {
	return &EmailService{}
}

func (s *EmailService) SendOtpEmail(email string, otp string) {

	// build payload
	payload := map[string]interface{}{
		"from": map[string]string{
			"email": "no-reply@project2morrow.com",
			"name":  "project 2morrow Software LTD",
		},
		"to": []map[string]string{
			{"email": email},
		},
		"subject": "Your verification code",
		"text":    fmt.Sprintf("Your OTP is %s. It expires in 10 minutes.", otp),
	}

	bodyBytes, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
		return
	}

	req, err := http.NewRequest("POST", "https://send.api.mailtrap.io/api/send", bytes.NewReader(bodyBytes))
	if err != nil {
		fmt.Println(err)
		return
	}

	// read API token from env; fallback to placeholder
	token := os.Getenv("MAILTRAP_API_TOKEN")
	if token == "" {
		token = "<YOUR_API_TOKEN>"
	}
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	respBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("mailtrap response:", string(respBody))
	fmt.Println("sent otp:", otp)
}
