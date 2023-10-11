package env

import (
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
	"os"
)

type Config struct {
	AppPort               string
	Key                   string
	FacebookAppSecret     string
	AppleSecretApi        string
	PublisherToken        string
	OffersFilePath        string
	PlayerDatasetFilePath string
	AwardPublisherUrl     string
	AssetUploadGatewayUrl string
	ReportingApiUrl       string
}

func New() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
	return &Config{
		AppPort:               os.Getenv("APP_PORT"),
		Key:                   os.Getenv("KEY"),
		FacebookAppSecret:     os.Getenv("FACEBOOK_APP_SECRET"),
		AppleSecretApi:        os.Getenv("APPLE_SECRET_API"),
		PublisherToken:        os.Getenv("PUBLISHER_TOKEN"),
		OffersFilePath:        os.Getenv("OFFERS_FILE_PATH"),
		PlayerDatasetFilePath: os.Getenv("PLAYER_DATASET_FILE_PATH"),
		AwardPublisherUrl:     os.Getenv("AWARD_PUBLISHER_URL"),
		AssetUploadGatewayUrl: os.Getenv("ASSET_UPLOAD_GATEWAY_URL"),
		ReportingApiUrl:       os.Getenv("REPORTING_API_URL"),
	}
}
