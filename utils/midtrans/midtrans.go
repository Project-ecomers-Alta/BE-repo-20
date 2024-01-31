package midtrans

import (
	"BE-REPO-20/app/configs"
	"BE-REPO-20/features/order"
	"errors"
	"fmt"
	"strconv"

	mid "github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type MidtransInterface interface {
	Order(data order.OrderCore) (*order.OrderCore, error)
	CancelOrder(orderId string) error
}

type midtrans struct {
	client coreapi.Client
	// environtment mid.EnvironmentType
}

func New() MidtransInterface {
	environment := mid.Sandbox
	var client coreapi.Client
	client.New(configs.MIDTRANS_SERVER_KEY, environment)

	return &midtrans{
		client: client,
	}
}

// Order implements MidtransInterface.
func (midtrans *midtrans) Order(data order.OrderCore) (*order.OrderCore, error) {
	req := new(coreapi.ChargeReq)
	// uuid := uuid.New()
	fmt.Println("order id :", strconv.Itoa(int(data.Id)))
	req.TransactionDetails = mid.TransactionDetails{
		OrderID: strconv.Itoa(int(data.Id)),
		// OrderID:  uuid.String(),
		GrossAmt: int64(data.Total),
	}

	switch data.PaymentMethod {
	case "BCA":
		req.PaymentType = coreapi.PaymentTypeBankTransfer
		req.BankTransfer = &coreapi.BankTransferDetails{
			Bank: mid.BankBca,
		}
	case "BNI":
		req.PaymentType = coreapi.PaymentTypeBankTransfer
		req.BankTransfer = &coreapi.BankTransferDetails{
			Bank: mid.BankBni,
		}
	case "BRI":
		req.PaymentType = coreapi.PaymentTypeBankTransfer
		req.BankTransfer = &coreapi.BankTransferDetails{
			Bank: mid.BankBri,
		}

	default:
		return nil, errors.New("payment not support")

	}

	res, err := midtrans.client.ChargeTransaction(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != "201" {
		return nil, errors.New(res.StatusMessage)
	}

	// response
	VaNumb, _ := (strconv.Atoi(res.VaNumbers[0].VANumber))
	fmt.Println("va number: ", VaNumb)
	data.VirtualAcc = uint(VaNumb)
	data.PaymentMethod = res.VaNumbers[0].Bank
	data.Status = res.TransactionStatus
	data.TransactionTime = res.TransactionTime

	// data.VaNumber, _ = strconv.Atoi(res.VaNumbers[0].VANumber)
	// data.PaymentType = res.PaymentType
	// data.Status = res.TransactionStatus
	// data.Payment.StatusCode = res.StatusCode
	// data.Payment.StatusMessage = res.StatusMessage
	// data.Payment.TransactionId = res.TransactionID
	// data.Payment.Currency = res.Currency
	// data.Payment.TransactionTime = res.TransactionTime
	// data.Payment.FraudStatus = res.FraudStatus
	// data.Payment.ExpiredAt = res.ExpiryTime

	return &data, nil
}

// CancelOrder implements MidtransInterface.
func (midtrans *midtrans) CancelOrder(orderId string) error {
	panic("unimplemented")
}
