package paymentcode

import (
	"context"
	"joshua-golang-training-beginner/errors"
	mockPayment "joshua-golang-training-beginner/mock_payment"
	payment "joshua-golang-training-beginner/payment"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func getMockPaymentCode() *payment.PaymentCode {
	return &payment.PaymentCode{
		PaymentCode: "test-payment",
		Name:        "tester",
	}
}

func TestServiceCreateSuccess(t *testing.T) {
	// initialize gomock
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPC := getMockPaymentCode()

	mockPaymentCodeRepo := mockPayment.NewMockIPaymentCodeRepository(ctrl)
	mockPaymentCodeService := mockPayment.NewMockIPaymentCodeService(ctrl)

	mockPaymentCodeRepo.EXPECT().Create(
		context.Background(),
		mockPC,
	).Return(nil)

	mockPaymentCodeService.EXPECT().Create(
		context.Background(),
		mockPC,
	).Return(nil)

	err := mockPaymentCodeService.Create(context.Background(), mockPC)
	require.NoError(t, err)
	err = mockPaymentCodeRepo.Create(context.Background(), mockPC)
	require.NoError(t, err)
}

func TestServiceCreateFail(t *testing.T) {
	// initialize gomock
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPC := getMockPaymentCode()
	mockPC.Name = ""

	mockPaymentCodeRepo := mockPayment.NewMockIPaymentCodeRepository(ctrl)
	mockPaymentCodeService := mockPayment.NewMockIPaymentCodeService(ctrl)

	expectedErr := errors.ValidationError("Payment code or name should not be empty.")

	mockPaymentCodeRepo.EXPECT().Create(
		context.Background(),
		mockPC,
	).Return(expectedErr)

	mockPaymentCodeService.EXPECT().Create(
		context.Background(),
		mockPC,
	).Return(expectedErr)

	err := mockPaymentCodeService.Create(context.Background(), mockPC)
	require.Error(t, err)
	err = mockPaymentCodeRepo.Create(context.Background(), mockPC)
	require.Error(t, err)
}

func TestServiceGetByIDSuccess(t *testing.T) {
	// initialize gomock
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expectedPC := payment.PaymentCode{
		PaymentCode:    "test-payment",
		Name:           "tester",
		ID:             uuid.New(),
		Status:         "ACTIVE",
		CreatedTime:    time.Now(),
		UpdatedTime:    time.Now(),
		ExpirationDate: time.Now().AddDate(50, 0, 0),
	}

	mockPaymentCodeRepo := mockPayment.NewMockIPaymentCodeRepository(ctrl)
	mockPaymentCodeService := mockPayment.NewMockIPaymentCodeService(ctrl)

	mockPaymentCodeRepo.EXPECT().GetByID(
		context.Background(),
		expectedPC.ID.String(),
	).Return(expectedPC, nil)

	mockPaymentCodeService.EXPECT().GetByID(
		context.Background(),
		expectedPC.ID.String(),
	).Return(expectedPC, nil)

	res, err := mockPaymentCodeService.GetByID(context.Background(), expectedPC.ID.String())
	require.NoError(t, err)
	require.Equal(t, expectedPC, res)
	res, err = mockPaymentCodeRepo.GetByID(context.Background(), expectedPC.ID.String())
	require.NoError(t, err)
	require.Equal(t, expectedPC, res)
}

func TestServiceGetByIDFail(t *testing.T) {
	// initialize gomock
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expectedPC := payment.PaymentCode{}

	mockPaymentCodeRepo := mockPayment.NewMockIPaymentCodeRepository(ctrl)
	mockPaymentCodeService := mockPayment.NewMockIPaymentCodeService(ctrl)

	expectedErr := errors.DataNotFoundError("Payment Code ID Not Found")

	mockPaymentCodeRepo.EXPECT().GetByID(
		context.Background(),
		expectedPC.ID.String(),
	).Return(expectedPC, expectedErr)

	mockPaymentCodeService.EXPECT().GetByID(
		context.Background(),
		expectedPC.ID.String(),
	).Return(expectedPC, expectedErr)

	res, err := mockPaymentCodeService.GetByID(context.Background(), expectedPC.ID.String())
	require.Error(t, err)
	require.Empty(t, res)
	res, err = mockPaymentCodeRepo.GetByID(context.Background(), expectedPC.ID.String())
	require.Error(t, err)
	require.Empty(t, res)
}
