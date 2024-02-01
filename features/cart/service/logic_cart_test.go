package service

import (
	auth "BE-REPO-20/features/cart"
	"BE-REPO-20/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSelectAllCart(t *testing.T) {
	repo := new(mocks.CartData)
	service := NewCart(repo)

	expectedCarts := []auth.CartCore{
		{ID: 1,
			UserID:    1,
			ProductID: 1,
			Quantity:  2}}

	repo.On("SelectAllCart", mock.Anything).Return(expectedCarts, nil).Once()

	carts, err := service.SelectAllCart(1)

	assert.NoError(t, err)
	assert.Equal(t, expectedCarts, carts)

	repo.AssertExpectations(t)
}

func TestDeleteCarts(t *testing.T) {
	repo := new(mocks.CartData)
	srv := NewCart(repo)

	ids := []uint{1, 2, 3}

	repo.On("DeleteCarts", mock.Anything).Return(nil).Once()

	err := srv.DeleteCarts(ids)

	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestCreateCart(t *testing.T) {
	repo := new(mocks.CartData)
	srv := NewCart(repo)

	input := auth.CartCore{
		UserID:    1,
		ProductID: 1,
		Quantity:  2}

	repo.On("CreateCart", mock.Anything).Return(nil).Once()

	err := srv.CreateCart(input)

	assert.NoError(t, err)
	repo.AssertExpectations(t)
}
