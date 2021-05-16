package main

import (
	db "joshua-golang-training-beginner/db"
	payment "joshua-golang-training-beginner/payment"
	paymentcode "joshua-golang-training-beginner/paymentcode"
	postgres "joshua-golang-training-beginner/postgres"
	"joshua-golang-training-beginner/rest"

	echo "github.com/labstack/echo/v4"
)

// Global variable for services to be passed to REST handlers
var (
	paymentCodeService payment.IPaymentCodeService
)

func main() {
	// Prepare DB
	db := db.InitDB()

	// Prepare repos
	paymentCodeRepo := postgres.NewPaymentCodeRepository(db)

	// Prepare services
	paymentCodeService = paymentcode.NewService(
		paymentCodeRepo,
	)

	// Prepare REST
	e := echo.New()
	registerControllers(e)
	e.Logger.Fatal(e.Start(":9091"))
}

// REST handlers
func registerControllers(e *echo.Echo) {
	rest.InitPaymentCodeHandler(e, paymentCodeService)
	rest.InitServerHandler(e)
}
