package httpresponse

import (
	"net/http"

	"github.com/go-mongoDB-example-github/models"
	"github.com/labstack/echo"
)

func InternalServerError(c echo.Context) error {
	return c.JSON(http.StatusInternalServerError, models.Error{
		Message: "Internal server error",
	})
}

func RecordNotFound(c echo.Context) error {
	return c.JSON(http.StatusNotFound, models.Error{
		Message: "Data not found",
	})
}

func BadRequest(c echo.Context, message string) error {
	if message == "" {
		message = "Invalid payload"
	}
	return c.JSON(http.StatusNotFound, models.Error{
		Message: message,
	})
}
