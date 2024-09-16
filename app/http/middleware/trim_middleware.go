package appMiddleware

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func TrimMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Read the request body
		body, err := io.ReadAll(c.Request().Body)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		// Restore the request body for later use
		c.Request().Body = io.NopCloser(bytes.NewBuffer(body))

		// Parse the request body into a map
		var requestData map[string]interface{}
		if err := json.Unmarshal(body, &requestData); err != nil {
			return next(c) // If it's not JSON, skip the middleware
		}

		// Trim all string fields
		for key, value := range requestData {
			if str, ok := value.(string); ok {
				requestData[key] = strings.TrimSpace(str)
			}
		}

		// Marshal the modified request data back into JSON
		modifiedBody, err := json.Marshal(requestData)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		// Replace the request body with the modified data
		c.Request().Body = io.NopCloser(bytes.NewBuffer(modifiedBody))

		return next(c)
	}
}
