package utilsjwt

import (
	"time"

	"loan_apps/config/setting"

	"loan_apps/scripts/utils"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var jwtKey = setting.JwtSecretKey
var jwtAud = setting.JwtAudience
var jwtIss = setting.JwtIssuer
var jwtAlg = jwt.SigningMethodHS256
var jwtRefreshExpired = setting.JwtRefreshExpiredPerDay
var jwtExpired = setting.JwtExpiredPerSession

// CustomClaims mengimplementasikan jwt.Claims
type CustomClaims struct {
	UserRole string `json:"user_role"`
	DeviceID string `json:"device_id"`
	jwt.RegisteredClaims
}

// GenerateAccessAndRefreshTokens generates access and refresh tokens
func GenerateAccessAndRefreshTokens(accessTokenID string, userAgent string) (string, string, error) {
	access, err := GenerateAccessToken(accessTokenID, userAgent)
	if err != nil {
		return "", "", err
	}

	refresh, err := GenerateRefreshAccessToken(accessTokenID, userAgent)
	if err != nil {
		return "", "", err
	}

	return access, refresh, nil
}

// GenerateAccessToken generates a JWT token based on the input token
func GenerateAccessToken(credential string, useragent string) (string, error) {

	jwtID := uuid.NewString()
	deviceID, err := utils.HashPassword(useragent)
	if err != nil {
		return "", err
	}

	claimster := CustomClaims{
		UserRole: "user",
		DeviceID: deviceID,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        jwtID, // Unique ID for tracking
			Issuer:    jwtIss,
			Subject:   credential,
			Audience:  jwt.ClaimStrings{jwtAud},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(jwtExpired) * time.Minute)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	return jwt.NewWithClaims(jwtAlg, claimster).SignedString([]byte(jwtKey))
}

// GenerateRefreshAccessToken generates a refresh JWT token based on the input token
func GenerateRefreshAccessToken(credential string, useragent string) (string, error) {
	jwtID := uuid.NewString()
	deviceID, err := utils.HashPassword(useragent)
	if err != nil {
		return "", err
	}

	// Create the Claims
	claimster := CustomClaims{
		UserRole: "user",
		DeviceID: deviceID,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        jwtID, // Unique ID for tracking
			Issuer:    jwtIss,
			Subject:   credential,
			Audience:  jwt.ClaimStrings{jwtAud},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(jwtRefreshExpired) * 24 * time.Hour)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	return jwt.NewWithClaims(jwtAlg, claimster).SignedString([]byte(jwtKey))
}

// ValidationAccessToken validates the JWT token string and returns the claims if valid
func ValidationAccessToken(tokenString string) (*CustomClaims, error) {
	// Create the Claims
	claims := &CustomClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}
	return claims, nil
}

// RefreshAccessToken validates the JWT token string and returns the claims if valid
func RefreshAccessToken(tokenString string, useragent string) (string, string, error) {
	// Create the Claims
	claims := &CustomClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})

	if err != nil || !token.Valid {
		return "", "", err
	}

	// Generate new access token
	newAccessToken, err := GenerateAccessToken(claims.RegisteredClaims.Subject, useragent)
	if err != nil {
		return "", "", err
	}

	// Optionally generate a new refresh token
	newRefreshToken, err := GenerateRefreshAccessToken(claims.RegisteredClaims.Subject, useragent)
	if err != nil {
		return "", "", err
	}
	return newAccessToken, newRefreshToken, nil
}
