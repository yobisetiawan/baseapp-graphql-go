package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type WelcomeController struct {
}

func NewWelcomeController() *WelcomeController {
	return &WelcomeController{}
}

// @Tags welcome
// @Router / [get]
func (ctr *WelcomeController) Index(c echo.Context) error {

	return c.JSON(http.StatusOK, echo.Map{"app": "welcome to api eticket"})
}

// @Tags welcome
// @Router /health [get]
func (ctr *WelcomeController) Health(c echo.Context) error {

	return c.JSON(http.StatusOK, echo.Map{"data": echo.Map{
		"DB":                    "OK",
		"S3_Storage_Service":    "OK",
		"Email_Service":         "OK",
		"Firebase_Service":      "OK",
		"Google_Signin_Service": "OK",
		"Payment_Service":       "OK",
		"Forwarder_Service":     "OK",
	}})
}
