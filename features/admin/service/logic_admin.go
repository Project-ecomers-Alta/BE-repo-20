package service

import "BE-REPO-20/features/admin"

type adminService struct {
	adminData admin.AdminDataInterface
}

func NewAdmin(adminData admin.AdminDataInterface) admin.AdminServiceInterface {
	return &adminService{
		adminData: adminData,
	}
}

// GetUserRoleById implements admin.AdminServiceInterface.
func (service *adminService) GetUserRoleById(userId int) (string, error) {
	return service.adminData.GetUserRoleById(userId)
}

// SelectAllUser implements admin.AdminServiceInterface.
func (service *adminService) SelectAllUser() ([]admin.AdminUserCore, error) {
	// logic
	// memanggil func yg ada di data layer
	results, err := service.adminData.SelectAllUser()
	return results, err

}
