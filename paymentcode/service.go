package paymentcode

import (
	"context"

	"joshua-golang-training-beginner/errors"
	payment "joshua-golang-training-beginner/payment"
)

type service struct {
	paymentCodeRepo payment.IPaymentCodeRepository
}

func NewService(
	paymentCodeRepo payment.IPaymentCodeRepository,
) payment.IPaymentCodeService {
	return &service{
		paymentCodeRepo: paymentCodeRepo,
	}
}

func (s service) GetByID(ctx context.Context, id string) (res payment.PaymentCode, err error) {
	res, err = s.paymentCodeRepo.GetByID(ctx, id)
	if err != nil {
		return
	}

	return
}

func (s service) Create(ctx context.Context, p *payment.PaymentCode) (err error) {
	p.Status = payment.PaymentCodeStatus.Active

	// Validate fields
	if p.PaymentCode == "" || p.Name == "" {
		err = errors.ValidationError("Payment code or name should not be empty.")
		return
	}

	err = s.paymentCodeRepo.Create(ctx, p)
	if err != nil {
		return
	}

	return
}
