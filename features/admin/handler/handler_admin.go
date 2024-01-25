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

// SearchUsersByQueryHandler handles the endpoint for searching users by query.
func (handler *AdminHandler) SearchUsersByQuery(c echo.Context) error {
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

	// Dapatkan query pencarian dari parameter URL
	query := c.QueryParam("search")

	// Panggil func di service layer untuk pencarian berdasarkan query
	results, errSearch := handler.adminService.SearchUserByQuery(query)
	if errSearch != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("Error searching data. "+errSearch.Error(), nil))
	}

	// Proses mapping dari core ke response
	usersResult := CoreToResponseList(results)

	return c.JSON(http.StatusOK, responses.WebResponse("Success searching data.", usersResult))
}

func (handler *AdminHandler) GetAllOrders(c echo.Context) error {
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
	results, errSelect := handler.adminService.SelectAllOrder()
	if errSelect != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("Error reading data. "+errSelect.Error(), nil))
	}

	// Proses mapping dari core ke response
	ordersResult := CoreToItemOrderResponseList(results)

	return c.JSON(http.StatusOK, responses.WebResponse("Success reading data.", ordersResult))
}

func (handler *AdminHandler) SearchOrderByQuery(c echo.Context) error {
	// Get user id from the token
	userId := middlewares.ExtractTokenUserId(c)

	// Get user role from the service layer
	userRole, err := handler.adminService.GetUserRoleById(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("Internal Server Error", nil))
	}

	// Check if the user role is 'admin'
	if userRole != "admin" {
		return c.JSON(http.StatusForbidden, responses.WebResponse("Forbidden - User is not an admin", nil))
	}

	// Get search query from the URL parameter
	query := c.QueryParam("search")

	// Call the service layer function for searching orders by query
	results, errSearch := handler.adminService.SearchOrderByQuery(query)
	if errSearch != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("Error searching data. "+errSearch.Error(), nil))
	}

	// Process mapping from core to response
	ordersResult := CoreToItemOrderResponseList(results)

	return c.JSON(http.StatusOK, responses.WebResponse("Success searching data.", ordersResult))
}
