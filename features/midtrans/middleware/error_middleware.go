package middleware

import (
	"BE-REPO-20/features/midtrans/helper"
	"BE-REPO-20/features/midtrans/web"
	"errors"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func ErrorHandle() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := next(c)

			if err != nil {
				if validationErrors(c, err) {
					return err
				}

				internalServerError(c, err)
			}

			return nil
		}
	}
}

func validationErrors(c echo.Context, err error) bool {
	if exception, ok := err.(validator.ValidationErrors); ok {
		var ve validator.ValidationErrors
		out := make([]web.ErrorResponse, len(ve))
		if errors.As(exception, &ve) {
			for _, fe := range ve {
				out = append(out, web.ErrorResponse{
					Field:   fe.Field(),
					Message: helper.MessageForTag(fe.Tag()),
				})
			}
		}
		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   out,
		}
		if err := c.JSON(http.StatusBadRequest, webResponse); err != nil {
			c.Error(err) // Set an error and stop middleware execution
			return true
		}
		c.Echo().Renderer = nil // Avoid rendering twice

		return true
	} else {
		return false
	}
}

func internalServerError(c echo.Context, err error) {
	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   err,
	}

	if err := c.JSON(http.StatusInternalServerError, webResponse); err != nil {
		c.Error(err) // Set an error and stop middleware execution
		return
	}
	c.Echo().Renderer = nil // Avoid rendering twice
}
