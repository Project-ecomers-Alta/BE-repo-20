package service

import (
	"BE-REPO-20/features/midtrans/helper"
	"BE-REPO-20/features/midtrans/web"
	"strconv"

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
	user_id := strconv.Itoa(request.UserId)

	// customer
	custAddress := &midtrans.CustomerAddress{
		FName:       "John",
		LName:       "Doe",
		Phone:       "081234567890",
		Address:     "Baker Street 97th",
		City:        "Jakarta",
		Postcode:    "16000",
		CountryCode: "IDN",
	}

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  "MID-User-" + user_id + "-" + request.OrderID,
			GrossAmt: request.Amount,
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName:    "John",
			LName:    "Doe",
			Email:    "john@doe.com",
			Phone:    "081234567890",
			BillAddr: custAddress,
			ShipAddr: custAddress,
		},
		EnabledPayments: snap.AllSnapPaymentType,
		Items: &[]midtrans.ItemDetails{
			{
				ID:    "Property-" + request.OrderID,
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
