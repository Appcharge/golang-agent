package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/types"
	"golang-agent/common/env"
	"golang-agent/pkg/analytics/config"
	"golang-agent/pkg/signature"
	"net/http"
)

type AnalyticsService interface {
	GetAnalytics(GetAnalyticsRequestData *config.GetAnalyticsRequest) (types.Object, error)
}

type AnalyticsServiceImpl struct {
	envConfig        *env.Config
	signatureService signature.SignatureHashingService
}

func NewAnalyticsService() *AnalyticsServiceImpl {
	envConfig := env.New()
	signatureService := signature.NewSignatureHashingService()
	return &AnalyticsServiceImpl{
		envConfig:        envConfig,
		signatureService: signatureService,
	}
}

func (s *AnalyticsServiceImpl) GetAnalytics(requestBody string) ([]config.GetAnalyticsResponse, error) {
	newSignature, err := s.signatureService.GenerateSignature(requestBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		"POST",
		s.envConfig.ReportingApiUrl+"/reporting/reports/analytics",
		bytes.NewBuffer([]byte(requestBody)),
	)
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return nil, err
	}
	req.Header.Set("signature", newSignature)
	req.Header.Set("x-publisher-token", s.envConfig.PublisherToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
		return nil, err
	}
	defer resp.Body.Close()
	var analyticsResponses []config.GetAnalyticsResponse

	parsingErr := json.NewDecoder(resp.Body).Decode(&analyticsResponses)
	if parsingErr != nil {
		fmt.Println("Error parsing HTTP response:", err)
		return nil, err
	}
	return analyticsResponses, nil
}
