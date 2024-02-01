package service

import (
	"BE-REPO-20/features/admin"
	"BE-REPO-20/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetUserRoleById(t *testing.T) {
	repo := new(mocks.AdminData)
	srv := NewAdmin(repo)
	expectedRole := "admin"
	repo.On("GetUserRoleById", mock.Anything).Return(expectedRole, nil).Once()

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
	repo.On("SelectAllUser", mock.Anything).Return(expectedUsers, nil).Once()

	result, err := srv.SelectAllUser()

	assert.NoError(t, err)
	assert.Equal(t, expectedUsers, result)
	repo.AssertExpectations(t)
}

func TestSearchUserByQuery(t *testing.T) {
	t.Run("With Results", func(t *testing.T) {
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
				UpdatedAt: time.Now(),
			},
		}
		repo.On("SearchUserByQuery", mock.Anything).Return(expectedUsers, nil).Once()

		result, err := srv.SearchUserByQuery("query")

		assert.NoError(t, err)
		assert.Equal(t, expectedUsers, result)
		repo.AssertExpectations(t)
	})

	t.Run("Without Results", func(t *testing.T) {
		repo := new(mocks.AdminData)
		srv := NewAdmin(repo)

		repo.On("SearchUserByQuery", mock.Anything).Return([]admin.AdminUserCore{}, nil).Once()

		results, err := srv.SearchUserByQuery("nonexistent")

		assert.NoError(t, err)
		assert.Empty(t, results)
		repo.AssertExpectations(t)
	})
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
			ID:            1,
			UserID:        1,
			Address:       "Depok",
			PaymentMethod: "bank_transfer",
			Status:        "pending",
			Invoice:       "INV-001",
			Total:         2000000,
			VirtualAcc:    "VA-001",
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now()}}}
	repo.On("SelectAllOrder", mock.Anything).Return(expectedOrders, nil).Once()

	result, err := srv.SelectAllOrder()

	assert.NoError(t, err)
	assert.Equal(t, expectedOrders, result)
	repo.AssertExpectations(t)
}

func TestSearchOrderByQuery(t *testing.T) {
	t.Run("With Results", func(t *testing.T) {
		repo := new(mocks.AdminData)
		srv := NewAdmin(repo)
		expectedOrders := []admin.AdminItemOrderCore{
			{
				ID:           1,
				OrderID:      1,
				ProductName:  "Sepatu Nike",
				ProductPrice: 1000000,
				Quantity:     2,
				SubTotal:     2000000,
				CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
				Order: admin.AdminOrderCore{
					ID:            1,
					UserID:        1,
					Address:       "Depok",
					PaymentMethod: "bank_transfer",
					Status:        "pending",
					Invoice:       "INV-001",
					Total:         200,
					VirtualAcc:    "VA-001",
					CreatedAt:     time.Now(),
					UpdatedAt:     time.Now(),
				},
			},
		}
		repo.On("SearchOrderByQuery", mock.Anything).Return(expectedOrders, nil).Once()

		result, err := srv.SearchOrderByQuery("query")

		assert.NoError(t, err)
		assert.Equal(t, expectedOrders, result)
		repo.AssertExpectations(t)
	})

	t.Run("Without Results", func(t *testing.T) {
		repo := new(mocks.AdminData)
		srv := NewAdmin(repo)

		repo.On("SearchOrderByQuery", mock.Anything).Return([]admin.AdminItemOrderCore{}, nil).Once()

		results, err := srv.SearchOrderByQuery("nonexistent")

		assert.NoError(t, err)
		assert.Empty(t, results)
		repo.AssertExpectations(t)
	})
}
