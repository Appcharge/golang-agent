package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"go/types"
	"golang-agent/common/env"
	"golang-agent/pkg/player/config"
	"golang-agent/pkg/signature"
	"io/ioutil"
	"net/http"
	"os"
)

type PlayerService interface {
	SyncInfo(PlayerInfoSyncRequest *config.PlayerInfoSyncRequest) (types.Object, error)
	UpdateBalance(PlayerInfoSyncRequest *config.PlayerInfoSyncRequest) (types.Object, error)
}

type PlayerServiceImpl struct {
	envConfig        *env.Config
	signatureService signature.SignatureHashingService
}

func NewPlayerService() *PlayerServiceImpl {
	envConfig := env.New()
	signatureService := signature.NewSignatureHashingService()
	return &PlayerServiceImpl{
		envConfig:        envConfig,
		signatureService: signatureService,
	}
}

func (s *PlayerServiceImpl) SyncInfo(playerId string) (*config.PlayerData, error) {
	playerData, err := s.LoadPlayerData(playerId)
	if err != nil {
		return nil, err
	}
	return playerData, nil
}

func (s *PlayerServiceImpl) UpdateBalance(requestBody string) (*config.PlayerUpdateBalanceResponse, error) {
	newSignature, err := s.signatureService.GenerateSignature(requestBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", s.envConfig.AwardPublisherUrl, bytes.NewBuffer([]byte(requestBody)))
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
	var updateBalanceResponse config.PlayerUpdateBalanceResponse

	parsingErr := json.NewDecoder(resp.Body).Decode(&updateBalanceResponse)
	if parsingErr != nil {
		fmt.Println("Error parsing HTTP response:", err)
		return nil, err
	}
	return &updateBalanceResponse, nil
}

func (s *PlayerServiceImpl) LoadPlayerData(playerId string) (*config.PlayerData, error) {
	playerDataFile, err := os.Open(s.envConfig.PlayerDatasetFilePath)
	if err != nil {
		return nil, errors.New("Failed to load player data!")
	}
	defer playerDataFile.Close()

	content, err := ioutil.ReadAll(playerDataFile)
	if err != nil {
		return nil, errors.New("Failed to load player data!")
	}

	var playerDataSchema config.PlayerDataFile
	err = json.Unmarshal(content, &playerDataSchema)
	if err != nil {
		return nil, errors.New("Failed to load player data!")
	}
	return &playerDataSchema.Player, nil
}
