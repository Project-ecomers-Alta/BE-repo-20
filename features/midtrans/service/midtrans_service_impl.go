package service

import (
	"BE-REPO-20/features/midtrans/helper"
	"BE-REPO-20/features/midtrans/web"
	"fmt"
	"strconv"

	// _dataOrder "BE-REPO-20/features/order/data"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type MidtransServiceImpl struct {
	Validate *validator.Validate
}

func NewMidtransServiceImpl(validate *validator.Validate) *MidtransServiceImpl {
	return &MidtransServiceImpl{
		Validate: validate,
	}
}

// var db *gorm.DB

func (service *MidtransServiceImpl) CreateEcho(c echo.Context, request web.MidtransRequest) web.MidtransResponse {
	err := service.Validate.Struct(request)
	if err != nil {
		helper.PanicIfError(err)
	}
	midtranServerKey := "SB-Mid-server-XGdfn0Un08oUioFXuEZuTkb-"
	// request midtrans
	var snapClient = snap.Client{}
	// snapClient.New(os.Getenv("MIDTRANS_SERVER_KEY"), midtrans.Sandbox)
	snapClient.New(midtranServerKey, midtrans.Sandbox)

	// user id
	// user_id := strconv.Itoa(request.UserId)
	// orderData := _dataOrder.NewOrder(db)
	fmt.Println(request.OrderID)
	// order, _ := order.OrderDataInterface.GetOrder(orderData, uint(request.OrderID))
	// var orderGorm _dataOrder.Order
	// order := db.Preload("ItemOrders").Preload("User").First(&orderGorm, "id = ?", uint(request.OrderID))
	// if order.Error != nil {
	// 	helper.PanicIfError(err)
	// }
	// customer
	custAddress := &midtrans.CustomerAddress{
		// FName:       "John",
		// LName:       "Doe",
		// Phone:       "081234567890",
		// Address:     "Baker Street 97th",
		// City:        "Jakarta",
		// Postcode:    "16000",
		// CountryCode: "IDN",

		// from GetOrder
		FName:       request.FName,
		LName:       request.LName,
		Phone:       request.Phone,
		Address:     request.Address,
		City:        request.City,
		Postcode:    "16000",
		CountryCode: "IDN",
	}

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			// OrderID:  "MID-User-" + user_id + "-" + request.OrderID,
			OrderID:  strconv.Itoa(int(request.OrderID)),
			GrossAmt: request.Amount,
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			// FName:    "John",
			FName:    request.FName,
			LName:    "",
			Email:    request.Email,
			Phone:    request.Phone,
			BillAddr: custAddress,
			ShipAddr: custAddress,
		},
		EnabledPayments: snap.AllSnapPaymentType,
		Items: &[]midtrans.ItemDetails{
			{
				ID:    "Property-" + strconv.Itoa(int(request.OrderID)),
				Qty:   1,
				Price: request.Amount,
				Name:  request.ItemName,
			},
		},
	}

	response, errSnap := snapClient.CreateTransaction(req)
	if errSnap != nil {
		helper.PanicIfError(errSnap.GetRawError())
	}

	midtransReponse := web.MidtransResponse{
		Token:       response.Token,
		RedirectUrl: response.RedirectURL,
	}

	return midtransReponse
}
