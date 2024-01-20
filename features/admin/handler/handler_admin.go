package handler

import (
	"BE-REPO-20/app/middlewares"
	"BE-REPO-20/features/admin"
	"BE-REPO-20/utils/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AdminHandler struct {
	adminService admin.AdminServiceInterface
}

func NewAdmin(service admin.AdminServiceInterface) *AdminHandler {
	return &AdminHandler{
		adminService: service,
	}
}

func (handler *AdminHandler) GetAllUsers(c echo.Context) error {
	// Dapatkan user id dari token
	userId := middlewares.ExtractTokenUserId(c)

	// Pergi ke service layer untuk mendapatkan role pengguna
	userRole, err := handler.adminService.GetUserRoleById(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("Internal Server Error", nil))
	}

	// Periksa apakah peran pengguna adalah 'admin'
	if userRole != "admin" {
		return c.JSON(http.StatusForbidden, responses.WebResponse("Forbidden - User is not an admin", nil))
	}

	// Jika peran pengguna adalah 'admin', panggil func di service layer
	results, errSelect := handler.adminService.SelectAllUser()
	if errSelect != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("Error reading data. "+errSelect.Error(), nil))
	}

	// Proses mapping dari core ke response
	usersResult := CoreToResponseList(results)

	return c.JSON(http.StatusOK, responses.WebResponse("Success reading data.", usersResult))
}
