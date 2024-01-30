package service

import (
	"BE-REPO-20/features/admin"
	"BE-REPO-20/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetUserRoleById(t *testing.T) {
	repo := new(mocks.AdminData)
	srv := NewAdmin(repo)
	expectedRole := "admin"
	repo.On("GetUserRoleById", 1).Return(expectedRole, nil)

	result, err := srv.GetUserRoleById(1)

	assert.NoError(t, err)
	assert.Equal(t, expectedRole, result)
	repo.AssertExpectations(t)
}

func TestSelectAllUser(t *testing.T) {
	repo := new(mocks.AdminData)
	srv := NewAdmin(repo)
	expectedUsers := []admin.AdminUserCore{
		{ID: 1,
			FullName:  "alta",
			UserName:  "alta academy",
			Email:     "alta@mail.com",
			Domicile:  "Jakarta",
			Password:  "password",
			Role:      "user",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now()}}
	repo.On("SelectAllUser").Return(expectedUsers, nil)

	result, err := srv.SelectAllUser()

	assert.NoError(t, err)
	assert.Equal(t, expectedUsers, result)
	repo.AssertExpectations(t)
}

func TestSearchUserByQuery(t *testing.T) {
	repo := new(mocks.AdminData)
	srv := NewAdmin(repo)
	expectedUsers := []admin.AdminUserCore{
		{ID: 1,
			FullName:  "alta",
			UserName:  "alta academy",
			Email:     "alta@mail.com",
			Domicile:  "Jakarta",
			Password:  "password",
			Role:      "user",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now()}}
	repo.On("SearchUserByQuery", "query").Return(expectedUsers, nil)

	result, err := srv.SearchUserByQuery("query")

	assert.NoError(t, err)
	assert.Equal(t, expectedUsers, result)
	repo.AssertExpectations(t)
}

func TestSelectAllOrder(t *testing.T) {
	repo := new(mocks.AdminData)
	srv := NewAdmin(repo)
	expectedOrders := []admin.AdminItemOrderCore{{
		ID:           1,
		OrderID:      1,
		ProductName:  "Sepatu Nike",
		ProductPrice: 1000000,
		Quantity:     2,
		SubTotal:     2000000,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		Order: admin.AdminOrderCore{
			ID:         1,
			UserID:     1,
			Address:    "Depok",
			CreditCard: "1234567890123456",
			Status:     "pending",
			Invoice:    "INV-001",
			Total:      2000000,
			VirtualAcc: "VA-001",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now()}}}
	repo.On("SelectAllOrder").Return(expectedOrders, nil)

	result, err := srv.SelectAllOrder()

	assert.NoError(t, err)
	assert.Equal(t, expectedOrders, result)
	repo.AssertExpectations(t)
}

func TestSearchOrderByQuery(t *testing.T) {
	repo := new(mocks.AdminData)
	srv := NewAdmin(repo)
	expectedOrders := []admin.AdminItemOrderCore{
		{ID: 1,
			OrderID:      1,
			ProductName:  "Sepatu Nike",
			ProductPrice: 1000000,
			Quantity:     2,
			SubTotal:     2000000,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			Order: admin.AdminOrderCore{
				ID:         1,
				UserID:     1,
				Address:    "Depok",
				CreditCard: "1234567890123456",
				Status:     "pending",
				Invoice:    "INV-001",
				Total:      200,
				VirtualAcc: "VA-001",
				CreatedAt:  time.Now(),
				UpdatedAt:  time.Now()}}}
	repo.On("SearchOrderByQuery", "query").Return(expectedOrders, nil)

	result, err := srv.SearchOrderByQuery("query")

	assert.NoError(t, err)
	assert.Equal(t, expectedOrders, result)
	repo.AssertExpectations(t)
}

func TestSearchUserByQuery_NoResults(t *testing.T) {

	repo := new(mocks.AdminData)
	srv := NewAdmin(repo)

	repo.On("SearchUserByQuery", "nonexistent").Return([]admin.AdminUserCore{}, nil)

	results, err := srv.SearchUserByQuery("nonexistent")

	assert.NoError(t, err)
	assert.Empty(t, results)
	repo.AssertExpectations(t)
}

func TestSearchOrderByQuery_NoResults(t *testing.T) {

	repo := new(mocks.AdminData)
	srv := NewAdmin(repo)

	repo.On("SearchOrderByQuery", "nonexistent").Return([]admin.AdminItemOrderCore{}, nil)

	results, err := srv.SearchOrderByQuery("nonexistent")

	assert.NoError(t, err)
	assert.Empty(t, results)
	repo.AssertExpectations(t)
}
