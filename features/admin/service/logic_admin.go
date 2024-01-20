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

// SearchUserByQuery implements admin.AdminServiceInterface.
func (service *adminService) SearchUserByQuery(query string) ([]admin.AdminUserCore, error) {
	// Panggil fungsi pencarian dari lapisan data (misalnya adminQuery)
	results, err := service.adminData.SearchUserByQuery(query)
	if err != nil {
		return nil, err
	}

	// Jika tidak ditemukan pengguna, kembalikan array kosong
	if len(results) == 0 {
		return []admin.AdminUserCore{}, nil
	}

	return results, nil
}
