package utils

import (
	"baseapp/app/configs"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func JWTGenerateToken(dt jwt.MapClaims, secretType string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, dt)
	secretKey := configs.AppConfig.JWTSecret
	if secretType == "refresh" {
		secretKey = secretKey + configs.AppConfig.JWTRefreshSecret
	}

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", echo.NewHTTPError(http.StatusInternalServerError, "Unable generatye token")
	}

	return tokenString, nil
}

func JWTCheckClaim(tokenString string, keyField string, secretType string) (string, error) {
	secretKey := configs.AppConfig.JWTSecret
	if secretType == "refresh" {
		secretKey = secretKey + configs.AppConfig.JWTRefreshSecret
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the token method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, echo.ErrUnauthorized
		}
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		return "", echo.NewHTTPError(http.StatusBadRequest, map[string]string{
			"message": "invalid or expired token",
		})
	}

	if token == nil {
		return "", echo.NewHTTPError(http.StatusBadRequest, map[string]string{
			"message": "invalid token",
		})
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", echo.NewHTTPError(http.StatusBadRequest, map[string]string{
			"message": "invalid token claim",
		})
	}

	userID, ok := claims[keyField].(string)
	if ok {
		return userID, nil
	}

	return "", echo.NewHTTPError(http.StatusBadRequest, map[string]string{
		"message": "invalid refresh token",
	})
}
