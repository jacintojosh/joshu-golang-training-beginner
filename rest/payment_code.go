package rest

import (
	"net/http"

	errors "joshua-golang-training-beginner/errors"
	payment "joshua-golang-training-beginner/payment"

	"github.com/labstack/echo/v4"
)

// Controller
type paymentCodeHandler struct {
	service payment.IPaymentCodeService
}

// Defines routes and initializes handlers
func InitPaymentCodeHandler(e *echo.Echo, service payment.IPaymentCodeService) {
	h := &paymentCodeHandler{service: service}

	e.POST("/payment-codes", h.Create)
	e.GET("/payment-codes/:id", h.Get)
}

// Handler function for route POST /payment-codes
func (h *paymentCodeHandler) Create(c echo.Context) (err error) {
	p := payment.PaymentCode{}

	if err = c.Bind(&p); err != nil {
		return
	}

	ctx := c.Request().Context()
	err = h.service.Create(ctx, &p)
	if err != nil {
		if err == errors.ValidationError("Payment code or name should not be empty.") {
			return c.JSON(errors.ValidationError.StatusCode(""), err)
		}
		return
	}

	return c.JSON(http.StatusCreated, p)
}

// Handler function for route GET /payment-codes/:id
func (h *paymentCodeHandler) Get(c echo.Context) (err error) {
	ctx := c.Request().Context()
	id := c.Param("id")
	res, err := h.service.GetByID(ctx, id)

	if err != nil {
		if err == errors.DataNotFoundError("Payment Code ID Not Found") {
			return c.JSON(errors.DataNotFoundError.StatusCode(""), err)
		}
		return
	}

	return c.JSON(http.StatusOK, res)
}
