package rest

import (
	"context"
	errors "joshua-golang-training-beginner/errors"
	payment "joshua-golang-training-beginner/payment"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	mockPayment "joshua-golang-training-beginner/mock_payment"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
)

func getMockPaymentCode() *payment.PaymentCode {
	return &payment.PaymentCode{
		PaymentCode: "test-payment",
		Name:        "tester",
	}
}

func TestCreateSuccess(t *testing.T) {
	e := echo.New()

	// initialize gomock
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPC := getMockPaymentCode()
	expectedPC := getMockPaymentCode()
	expectedPC.Status = "ACTIVE"

	mockPaymentCodeService := mockPayment.NewMockIPaymentCodeService(ctrl)
	mockPaymentCodeService.EXPECT().
		Create(
			context.Background(),
			mockPC,
		).
		Return(nil)

	reqBody := strings.NewReader(`{
		"payment_code": "test-payment",
		"name": "tester"
	}`)

	InitPaymentCodeHandler(e, mockPaymentCodeService)

	req := httptest.NewRequest("POST", "/payment-codes", reqBody)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	require.Equal(t, http.StatusCreated, rec.Code)
}

func TestCreateFail(t *testing.T) {
	e := echo.New()

	// initialize gomock
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPC := &payment.PaymentCode{
		PaymentCode: "test-payment",
	}

	err := errors.ValidationError("Payment code or name should not be empty.")

	mockPaymentCodeService := mockPayment.NewMockIPaymentCodeService(ctrl)
	mockPaymentCodeService.EXPECT().
		Create(
			context.Background(),
			mockPC,
		).
		Return(err)

	reqBody := strings.NewReader(`{
		"payment_code": "test-payment"
	}`)

	InitPaymentCodeHandler(e, mockPaymentCodeService)

	req := httptest.NewRequest("POST", "/payment-codes", reqBody)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	require.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestGetSuccess(t *testing.T) {
	e := echo.New()

	// initialize gomock
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	expectedPC := payment.PaymentCode{
		Status:      "ACTIVE",
		Name:        "tester",
		PaymentCode: "mock-code",
	}

	mockPaymentCodeService := mockPayment.NewMockIPaymentCodeService(ctrl)
	mockPaymentCodeService.EXPECT().
		GetByID(
			context.Background(),
			"mock-id",
		).
		Return(expectedPC, nil)

	reqBody := strings.NewReader(`{}`)

	InitPaymentCodeHandler(e, mockPaymentCodeService)

	req := httptest.NewRequest("GET", "/payment-codes/mock-id", reqBody)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	require.Equal(t, http.StatusOK, rec.Code)
}

func TestGetFail(t *testing.T) {
	e := echo.New()

	// initialize gomock
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	err := errors.DataNotFoundError("Payment Code ID Not Found")

	mockPaymentCodeService := mockPayment.NewMockIPaymentCodeService(ctrl)
	mockPaymentCodeService.EXPECT().
		GetByID(
			context.Background(),
			"fake-id",
		).
		Return(payment.PaymentCode{}, err)

	reqBody := strings.NewReader(`{}`)

	InitPaymentCodeHandler(e, mockPaymentCodeService)

	req := httptest.NewRequest("GET", "/payment-codes/fake-id", reqBody)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	require.Equal(t, http.StatusNotFound, rec.Code)
}

func TestInvalidUrl(t *testing.T) {
	e := echo.New()

	// initialize gomock
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPaymentCodeService := mockPayment.NewMockIPaymentCodeService(ctrl)

	reqBody := strings.NewReader(`{}`)

	InitPaymentCodeHandler(e, mockPaymentCodeService)

	req := httptest.NewRequest("GET", "/payment-coders", reqBody)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	require.Equal(t, http.StatusNotFound, rec.Code)
}
