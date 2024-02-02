package service

import (
	"BE-REPO-20/features/order"
	"BE-REPO-20/mocks"
	"errors"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPostOrder(t *testing.T) {
	repo := new(mocks.OrderData)
	srv := NewOrder(repo)

	userId := uint(1)
	input := order.OrderCore{
		Address:       "Jakarta",
		PaymentMethod: "BCA",
	}

	// invoice format (YYYYMMDD/NNNNN)
	today := time.Now()
	expectedInvoice := strconv.Itoa(today.Year()) +
		strconv.Itoa(int(today.Month())) +
		strconv.Itoa(today.Day()) + "/"

	// Subtest: Invalid User ID
	t.Run("Invalid User ID", func(t *testing.T) {
		_, err := srv.PostOrder(0, input)

		assert.Error(t, err)
		assert.Equal(t, "invalid id", err.Error())
	})

	// Subtest: Valid User ID
	t.Run("Valid User ID", func(t *testing.T) {
		repo.On("PostOrder", userId, mock.MatchedBy(func(input order.OrderCore) bool {
			return strings.HasPrefix(input.Invoice, expectedInvoice)
		})).Return(&order.OrderCore{}, nil).Once()

		_, err := srv.PostOrder(userId, input)
		assert.NoError(t, err)
		repo.AssertCalled(t, "PostOrder", userId, mock.MatchedBy(func(input order.OrderCore) bool {
			return strings.HasPrefix(input.Invoice, expectedInvoice)
		}))
	})
}

func TestGetOrders(t *testing.T) {
	repo := new(mocks.OrderData)
	srv := NewOrder(repo)

	userId := uint(1)
	expectedResult := []order.OrderCore{
		{
			Id:              1,
			UserID:          userId,
			Address:         "Jakarta",
			PaymentMethod:   "BCA",
			TransactionTime: "2024-02-01 22:07:08",
			Status:          "pending",
			Invoice:         "20240201/", // Misalnya invoice format YYYYMMDD
			Total:           100000000,
			VirtualAcc:      12344364,
		},
	}

	repo.On("GetOrders", userId).Return(expectedResult, nil).Once()

	result, err := srv.GetOrders(userId)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedResult, result)

	repo.AssertCalled(t, "GetOrders", userId)
}

func TestCancelOrder(t *testing.T) {
	repo := new(mocks.OrderData)
	srv := NewOrder(repo)

	userIdLogin := 1
	orderId := "12345"
	orderCore := order.OrderCore{Status: "pending"}

	repo.On("CancelOrder", userIdLogin, orderId, orderCore).Return(nil)

	err := srv.CancelOrder(userIdLogin, orderId, orderCore)

	assert.NoError(t, err)
	repo.AssertCalled(t, "CancelOrder", userIdLogin, orderId, orderCore)
}

func TestWebhoocksService(t *testing.T) {
	repo := new(mocks.OrderData)
	srv := orderService{repo}

	// Invalid Order ID
	t.Run("Invalid Order ID", func(t *testing.T) {
		webhoocksReq := order.OrderCore{Id: 0}
		expectedError := errors.New("invalid order id")

		repo.On("WebhoocksData", webhoocksReq).Return(nil).Once()

		err := srv.WebhoocksService(webhoocksReq)
		assert.EqualError(t, err, expectedError.Error(), "Expected an error for invalid order ID")
	})

	// Valid Order ID
	t.Run("Valid Order ID", func(t *testing.T) {
		webhoocksReq := order.OrderCore{Id: 1}

		repo.On("WebhoocksData", webhoocksReq).Return(nil).Once()

		err := srv.WebhoocksService(webhoocksReq)
		assert.NoError(t, err, "Expected no error for valid order ID")
	})
}
