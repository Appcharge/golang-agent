package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"go/types"
	"golang-agent/common/env"
	"golang-agent/pkg/offer/config"
	"golang-agent/pkg/signature"
	"io/ioutil"
	"net/http"
	"os"
)

type OfferService interface {
	CreateOffer(GetAnalyticsRequestData *config.CreateOfferRequest) (types.Object, error)
}

type OfferServiceImpl struct {
	envConfig        *env.Config
	signatureService signature.SignatureHashingService
}

func NewOfferService() *OfferServiceImpl {
	envConfig := env.New()
	signatureService := signature.NewSignatureHashingService()
	return &OfferServiceImpl{
		envConfig:        envConfig,
		signatureService: signatureService,
	}
}

func (s *OfferServiceImpl) CreateOffer(requestBody string) (*config.OfferResponse, error) {
	offersDataFile, err := os.Open(s.envConfig.OffersFilePath)
	if err != nil {
		return nil, errors.New("Failed to load offer data!")
	}
	defer offersDataFile.Close()

	byteContent, err := ioutil.ReadAll(offersDataFile)
	if err != nil {
		return nil, errors.New("Failed to load offer data!")
	}
	var offersDataObject map[string]interface{}
	json.Unmarshal(byteContent, &offersDataObject)

	jsonData, err := json.Marshal(offersDataObject["create"])
	if err != nil {
		fmt.Println("Error marshaling 'create' data:", err)
		return nil, err
	}
	newSignature, err := s.signatureService.GenerateSignature(string(jsonData))
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(
		"POST",
		s.envConfig.AssetUploadGatewayUrl+"/offering/offer/",
		bytes.NewBuffer(jsonData),
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
	fmt.Println(resp)
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
		return nil, err
	}
	defer resp.Body.Close()
	var createOfferResponse config.OfferResponse

	parsingErr := json.NewDecoder(resp.Body).Decode(&createOfferResponse)
	if parsingErr != nil {
		fmt.Println("Error parsing HTTP response:", err)
		return nil, err
	}
	return &createOfferResponse, nil
}

func (s *OfferServiceImpl) UpdateOffer(requestBody string) (*config.OfferResponse, error) {
	offersDataFile, err := os.Open(s.envConfig.OffersFilePath)
	if err != nil {
		return nil, errors.New("Failed to load offer data!")
	}
	defer offersDataFile.Close()

	byteContent, err := ioutil.ReadAll(offersDataFile)
	if err != nil {
		return nil, errors.New("Failed to load offer data!")
	}
	var offersDataObject map[string]interface{}
	json.Unmarshal(byteContent, &offersDataObject)

	jsonData, err := json.Marshal(offersDataObject["update"])
	if err != nil {
		fmt.Println("Error marshaling 'create' data:", err)
		return nil, err
	}
	newSignature, err := s.signatureService.GenerateSignature(string(jsonData))
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(
		"PUT",
		s.envConfig.AssetUploadGatewayUrl+"/offering/offer/",
		bytes.NewBuffer(jsonData),
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
	fmt.Println(resp)
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
		return nil, err
	}
	defer resp.Body.Close()
	var updateOfferResponse config.OfferResponse

	parsingErr := json.NewDecoder(resp.Body).Decode(&updateOfferResponse)
	if parsingErr != nil {
		fmt.Println("Error parsing HTTP response:", err)
		return nil, err
	}
	return &updateOfferResponse, nil
}
