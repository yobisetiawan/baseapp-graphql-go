package utils

import (
	"baseapp/app/database"
	"baseapp/app/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CurrentUser(c echo.Context, user *models.User) error {
	user_id, _ := ConvertToString(c.Get("user_id"))
	if user_id == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
	}

	database.DB.Model(models.User{}).Where("id = ?", user_id).First(&user)

	return nil
}
