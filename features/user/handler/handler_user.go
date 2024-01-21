package handler

import (
	"BE-REPO-20/app/middlewares"
	"BE-REPO-20/features/user"
	"BE-REPO-20/utils/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService user.UserServiceInterface
}

func New(service user.UserServiceInterface) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}

func (handler *UserHandler) SelectUser(c echo.Context) error {
	idJWT := middlewares.ExtractTokenUserId(c)
	if idJWT == 0 {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("unauthorized or jwt expired", nil))
	}

	data, err := handler.userService.SelectUser(idJWT)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error read data. "+err.Error(), nil))
	}

	result := CoreToResponse(*data)

	return c.JSON(http.StatusOK, responses.WebResponse("read success", result))
}

func (handler *UserHandler) Delete(c echo.Context) error {
	idJWT := middlewares.ExtractTokenUserId(c)
	if idJWT == 0 {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("unauthorized or jwt expired", nil))
	}
	err := handler.userService.Delete(idJWT)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error delete data. delete failed"+err.Error(), nil))
	}
	return c.JSON(http.StatusOK, responses.WebResponse("success delete data", nil))
}
