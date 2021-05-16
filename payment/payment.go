package payment

import (
	"context"
	"time"
)

type paymentCodeStatus struct {
	Active   string
	Inactive string
	Expired  string
}

var PaymentCodeStatus = paymentCodeStatus{
	Active:   "ACTIVE",
	Inactive: "INACTIVE",
	Expired:  "EXPIRED",
}

type (
	PaymentCode struct {
		ID             string    `json:"id"`
		Name           string    `json:"name"`
		PaymentCode    string    `json:"payment_code"`
		Status         string    `json:"status"`
		ExpirationDate time.Time `json:"expiration_date"`
		CreatedTime    time.Time `json:"created"`
		UpdatedTime    time.Time `json:"updated"`
	}
)

type IPaymentCodeRepository interface {
	Create(ctx context.Context, p *PaymentCode) error
	GetByID(ctx context.Context, id string) (p PaymentCode, err error)
}

type IPaymentCodeService interface {
	Create(ctx context.Context, p *PaymentCode) (err error)
	GetByID(ctx context.Context, id string) (p PaymentCode, err error)
}
