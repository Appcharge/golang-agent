package service

import (
	"encoding/json"
	"fmt"
	"go/types"
	"golang-agent/common/env"
	"golang-agent/pkg/order/config"
	"golang-agent/pkg/signature"
	"net/http"
	"strings"
)

type OrderService interface {
	GetOrders(getOrdersRequest *config.GetOrdersRequest) (types.Object, error)
}

type OrderServiceImpl struct {
	envConfig        *env.Config
	signatureService signature.SignatureHashingService
}

func NewOrderService() *OrderServiceImpl {
	envConfig := env.New()
	signatureService := signature.NewSignatureHashingService()
	return &OrderServiceImpl{
		envConfig:        envConfig,
		signatureService: signatureService,
	}
}

func (s *OrderServiceImpl) GetOrders(requestBody string) (*config.GetOrdersResponse, error) {
	newSignature, err := s.signatureService.GenerateSignature(requestBody)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(
		"POST",
		s.envConfig.ReportingApiUrl+"/reporting/reports/orders/",
		strings.NewReader(requestBody),
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
	var getOrdersResponse config.GetOrdersResponse

	parsingErr := json.NewDecoder(resp.Body).Decode(&getOrdersResponse)
	if parsingErr != nil {
		fmt.Println("Error parsing HTTP response:", err)
		return nil, err
	}
	return &getOrdersResponse, nil
}
