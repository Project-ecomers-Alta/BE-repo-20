package handler

import (
	"BE-REPO-20/app/middlewares"
	"BE-REPO-20/features/user"
	"BE-REPO-20/utils/responses"
	"fmt"
	"mime/multipart"
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

func (handler *UserHandler) UpdateUser(c echo.Context) error {
	// var fileType string
	var fileSize int64
	idJWT := middlewares.ExtractTokenUserId(c)
	if idJWT == 0 {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("unauthorized or jwt expired", nil))
	}

	var reqData = UserRequest{}
	errBind := c.Bind(&reqData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind, data not valid", nil))
	}
	userCore := RequestToCore(reqData)

	// upload img
	fileHeader, _ := c.FormFile("image_url")
	var file multipart.File
	if fileHeader != nil {
		openFileHeader, _ := fileHeader.Open()
		file = openFileHeader
	}
	// fileByte, _ := io.ReadAll(file)
	// fileType = http.DetectContentType(fileByte)
	// if fileType == "image/jpeg" || fileType == "image/png" || fileType == "image/webp" {
	// } else {
	// 	fmt.Println(fileType)
	// 	return c.JSON(http.StatusBadRequest, responses.WebResponse("error type data, data not valid", nil))
	// }

	fileSize = fileHeader.Size
	if fileSize >= 2000000 {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error size data, file size is too big", nil))
	}
	nameFile := fileHeader.Filename
	fmt.Println(nameFile)

	err := handler.userService.UpdateUser(idJWT, userCore, file, nameFile)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error update data. update failed"+err.Error(), nil))
	}
	return c.JSON(http.StatusOK, responses.WebResponse("success update data", nil))
}
