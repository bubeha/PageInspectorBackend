package auth

import (
	"os"
	"time"

	"github.com/bubeha/PageInspectorBackend/internal/infrastructure/config"
	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	jwt.RegisteredClaims
}

type TokenPairs struct {
	AccessToken  string        `json:"access_token"`
	RefreshToken string        `json:"refresh_token"`
	ExpiresAt    time.Duration `json:"expires_at"`
}

type TokenService interface {
	GeneratePairs(params any) (*TokenPairs, error)
}

type JWTTokenService struct {
	config *config.JWTConfig
}

func (ts *JWTTokenService) GeneratePairs(params any) (*TokenPairs, error) {
	privateKeyData, err := os.ReadFile(ts.config.PrivateKeyPath)
	if err != nil {
		return nil, err
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyData)

	if err != nil {
		return nil, err
	}

	publicToken := jwt.NewWithClaims(jwt.SigningMethodRS256, &CustomClaims{
		jwt.RegisteredClaims{
			Subject:   "User name",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ts.config.PublicTTL * time.Hour)),
		},
	})

	privateToken := jwt.NewWithClaims(jwt.SigningMethodRS256, &CustomClaims{
		jwt.RegisteredClaims{
			Subject:   "User name",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ts.config.PrivateTTL * time.Hour)),
		},
	})

	accessToken, err := publicToken.SignedString(privateKey)

	if err != nil {
		return nil, err
	}

	refreshToken, err := privateToken.SignedString(privateKey)

	if err != nil {
		return nil, err
	}

	return &TokenPairs{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    ts.config.PublicTTL,
	}, nil
}

func NewTokenService(cnf *config.JWTConfig) TokenService {
	return &JWTTokenService{config: cnf}
}
