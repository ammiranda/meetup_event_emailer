package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/ammiranda/meetup_emailer/ai"
	"github.com/ammiranda/meetup_emailer/config"
	"github.com/ammiranda/meetup_emailer/mailer"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environmental variables")
	}

	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	jsonFile, err := os.Open(config.JSONFilePath)
	if err != nil {
		log.Fatalf("Error opening JSON file: %v", err)
	}

	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatalf("Error reading JSON file: %v", err)
	}

	jsonDataString := string(jsonData)

	openaiClient := ai.NewOpenAI(context.Background(), config.OpenAIAPIKey)

	email, err := openaiClient.GenerateEmailBody(context.Background(), jsonDataString)
	if err != nil {
		log.Fatalf("Error generating email: %v", err)
	}

	mailer := mailer.NewMailer(config.SMTPHost, config.SMTPPort, config.SMTPUser, config.SMTPPassword, config.SMTPReceipents)

	err = mailer.SendHTMLEmail(context.Background(), "alexandermichaelmiranda@gmail.com", "Daily Networking Meetups", email)
	if err != nil {
		log.Fatalf("Error sending email: %v", err)
	}

	log.Println("Email sent successfully")
}
