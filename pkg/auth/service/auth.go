package service

import (
	"errors"
	"golang-agent/common/env"
	"golang-agent/pkg/auth/config"
	"golang-agent/pkg/auth/service/auth_method"
	_ "net/http"
)

type AuthenticationService interface {
	AuthenticatePlayer(authRequest *config.AuthenticationRequest) (*config.AuthenticationResponse, error)
}

type AuthenticationServiceImpl struct {
	FacebookAuth auth_method.FacebookAuth
	GoogleAuth   auth_method.GoogleAuth
	AppleAuth    auth_method.AppleAuth
}

func (a *AuthenticationServiceImpl) AuthenticatePlayer(authRequest *config.AuthenticationRequest) (*config.AuthenticationResponse, error) {
	var authResult config.AuthenticationResult

	envConfig := env.New()

	if authRequest == nil {
		return nil, errors.New("AuthenticationRequest is nil")
	}

	switch authMethod := authRequest.AuthMethod; authMethod {
	case "facebook":
		authResult = a.FacebookAuth.Authenticate(authRequest.AppID, *authRequest.Token, envConfig.FacebookAppSecret)
	case "google":
		authResult = a.GoogleAuth.Authenticate(authRequest.AppID, *authRequest.Token)
	case "apple":
		authResult = a.AppleAuth.Authenticate(authRequest.AppID, *authRequest.Token, envConfig.AppleSecretApi)
	case "userToken":
		authResult = config.AuthenticationResult{IsValid: true, UserID: *authRequest.Token}
	case "userPassword":
		authResult = config.AuthenticationResult{IsValid: true, UserID: *authRequest.UserName}
	default:
		return nil, errors.New("Unknown authentication method")
	}

	if authResult.IsValid {
		return createAuthResponse(&authResult), nil
	}

	return nil, errors.New("Authentication failed")
}

func createAuthResponse(authResult *config.AuthenticationResult) *config.AuthenticationResponse {
	return &config.AuthenticationResponse{
		Status:             "valid",
		PlayerProfileImage: "https://scontent.ftlv15-1.fna.fbcdn.net/v/t1.6435-9/39453230_281250465987441_6821580385961377792_n.jpg?_nc_cat=101&ccb=1-7&_nc_sid=09cbfe&_nc_ohc=EhNdfDjnT0MAX9Urj2h&_nc_ht=scontent.ftlv15-1.fna&oh=00_AfDz7cKzDCQC17o4L0i1ujpjilH11pTdfVyWPXMxzuOGxQ&oe=64C269BF",
		PublisherPlayerId:  authResult.UserID,
		PlayerName:         "Joe Dow",
		Segments:           []string{"seg1", "seg2"},
		Balances:           []config.ItemBalance{{Item: "diamonds", Amount: 15}},
	}
}
