package signature

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"golang-agent/common/env"
	"golang-agent/pkg/signature/config"
	"strconv"
	"strings"
	"time"
)

type SignatureHashingService interface {
	SignPayload(expectedSignatureHeader string, data string) (config.SignatureResponse, error)
	GenerateSignature(requestBody string) (string, error)
}

type SignatureHashingServiceImpl struct {
	envConfig *env.Config
}

func NewSignatureHashingService() *SignatureHashingServiceImpl {
	envConfig := env.New()
	return &SignatureHashingServiceImpl{
		envConfig: envConfig,
	}
}

func parseSignature(signatureString string) (*config.Signature, error) {
	parts := strings.Split(signatureString, ",")
	if len(parts) != 2 {
		return nil, errors.New("Invalid signature format")
	}

	tPart := strings.Split(parts[0], "=")
	v1Part := strings.Split(parts[1], "=")

	if len(tPart) != 2 || len(v1Part) != 2 {
		return nil, errors.New("Invalid signature format")
	}

	t, err := strconv.ParseInt(tPart[1], 10, 64)

	if err != nil {
		return nil, err
	}

	currentTime := time.Now().UnixMilli()
	elapsedTimeMinutes := (currentTime - t) / (1000 * 60)
	if elapsedTimeMinutes > 5 {
		return nil, errors.New("Timestamp is older than 5 minutes")
	}

	return &config.Signature{
		T:  tPart[1],
		V1: v1Part[1],
	}, nil
}

func compactStringJson(jsonBody string) string {
	return strings.ReplaceAll(strings.ReplaceAll(jsonBody, " ", ""), "\n", "")
}

func (s *SignatureHashingServiceImpl) SignPayload(signatureHeader string, data string) (config.SignatureResponse, error) {
	parsedSignature, err := parseSignature(signatureHeader)
	if err != nil {
		return config.SignatureResponse{}, err
	}

	preSignature := parsedSignature.T + "." + compactStringJson(data)

	hmacHash := hmac.New(sha256.New, []byte(s.envConfig.Key))
	_, err = hmacHash.Write([]byte(preSignature))
	if err != nil {
		return config.SignatureResponse{}, err
	}
	signature := hex.EncodeToString(hmacHash.Sum(nil))

	if signature != parsedSignature.V1 {
		return config.SignatureResponse{}, errors.New("Sinature is incorrect.")
	}

	return config.SignatureResponse{
		Signature: signature,
		V1:        parsedSignature.V1,
	}, nil
}

func (s *SignatureHashingServiceImpl) GenerateSignature(requestData string) (string, error) {
	currentTimeStamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	preSignature := currentTimeStamp + "." + compactStringJson(requestData)
	fmt.Println(preSignature)
	hmacHash := hmac.New(sha256.New, []byte(s.envConfig.Key))
	_, err := hmacHash.Write([]byte(compactStringJson(preSignature)))
	if err != nil {
		return "", err
	}
	signature := hex.EncodeToString(hmacHash.Sum(nil))
	fullSignature := "t=" + currentTimeStamp + "," + "v1=" + signature

	return fullSignature, nil
}
