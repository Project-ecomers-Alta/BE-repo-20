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

func (handler *UserHandler) UpdateUser(c echo.Context) error {
	idJWT := middlewares.ExtractTokenUserId(c)
	if idJWT == 0 {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("unauthorized or jwt expired", nil))
	}

	// upload img
	fileHeader, _ := c.FormFile("image_url")
	file, _ := fileHeader.Open()
	imgUrl, errUpload := middlewares.CloudinaryUpload(file)
	if errUpload != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error upload img", nil))
	}

	var reqData = UserRequest{}
	errBind := c.Bind(&reqData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind, data not valid", nil))
	}

	userCore := RequestToCore(reqData)
	userCore.Image = imgUrl.SecureURL
	// fmt.Println(userCore.Image)
	err := handler.userService.UpdateUser(idJWT, userCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error update data. update failed"+err.Error(), nil))
	}
	return c.JSON(http.StatusOK, responses.WebResponse("success update data", nil))
}
