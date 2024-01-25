package handler

import (
	"BE-REPO-20/app/middlewares"
	"BE-REPO-20/features/auth"
	"BE-REPO-20/utils/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	authService auth.AuthServiceInterface
}

func NewAuth(service auth.AuthServiceInterface) *AuthHandler {
	return &AuthHandler{
		authService: service,
	}
}

func (handler *AuthHandler) Register(c echo.Context) error {
	newUser := RegisterRequest{}
	errBind := c.Bind(&newUser)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid."+errBind.Error(), nil))
	}

	autCore := RequestToCore(newUser)

	_, token, errRegister := handler.authService.Register(autCore)
	if errRegister != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error insert data. insert failed"+errRegister.Error(), nil))
	}

	responseData := AuthRespon{
		UserName: newUser.UserName,
		Email:    newUser.Email,
		Domicile: newUser.Domicile,
		Role:     autCore.Role,
		Token:    token,
	}

	return c.JSON(http.StatusCreated, responses.WebResponse("insert success", responseData))
}

func (handler *AuthHandler) Login(c echo.Context) error {
	var reqData = LoginRequest{}
	errBind := c.Bind(&reqData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid"+errBind.Error(), nil))
	}

	result, token, err := handler.authService.Login(reqData.Email, reqData.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, responses.WebResponse(err.Error(), nil))
	}

	responseData := map[string]any{
		"id":        result.ID,
		"user_name": result.UserName,
		"email":     result.Email,
		"domicile":  result.Domicile,
		"role":      result.Role,
		"token":     token,
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success login.", responseData))
}

func (handler *AuthHandler) UpdatePassword(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)

	updatePassword := UpdatePasswordRequest{}
	errBind := c.Bind(&updatePassword)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid"+errBind.Error(), nil))
	}

	userCore := RequestToUpdatePassword(updatePassword)

	errUpdate := handler.authService.UptdatePassword(uint(userId), userCore)
	if errUpdate != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error editing data. "+errUpdate.Error(), nil))
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "Successful Operation",
	})
}
