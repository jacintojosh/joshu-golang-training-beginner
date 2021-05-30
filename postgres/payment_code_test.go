package postgres

import (
	"context"
	"database/sql"
	"joshua-golang-training-beginner/payment"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	_postgres "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/suite"
)

type paymentTestSuite struct {
	suite.Suite
	DSN    string
	DBConn *sql.DB
	DBName string
}

//https://pkg.go.dev/github.com/stretchr/testify/suite#pkg-overview
func TestSuitePayment(t *testing.T) {
	// temp db
	dsn := "user=postgres password=password dbname=golangtrainingdb host=localhost port=5432 sslmode=disable"

	paymentSuite := &paymentTestSuite{
		DSN: dsn,
	}

	suite.Run(t, paymentSuite)
}

func (s *paymentTestSuite) SetupSuite() {
	var err error
	s.DBConn, err = sql.Open("postgres", s.DSN)
	s.Require().NoError(err)
	err = s.DBConn.Ping()
	s.Require().NoError(err)
	s.Require().NoError(err)
}

//https://github.com/golang-migrate/migrate#use-in-your-go-project
func (s paymentTestSuite) BeforeTest(_, _ string) {
	driver, err := _postgres.WithInstance(s.DBConn, &_postgres.Config{})
	s.Require().NoError(err)

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", driver)
	s.Require().NoError(err)
	m.Up()
}

func (s paymentTestSuite) AfterTest(_, _ string) {
	driver, err := _postgres.WithInstance(s.DBConn, &_postgres.Config{})
	s.Require().NoError(err)

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", driver)
	s.Require().NoError(err)
	m.Down()
}

func getMockPaymentCode() *payment.PaymentCode {
	return &payment.PaymentCode{
		PaymentCode: "test-payment",
		Name:        "tester",
		Status:      "ACTIVE",
	}
}

func getInvalidMockPaymentCode() *payment.PaymentCode {
	return &payment.PaymentCode{
		PaymentCode: "test-payment",
		Status:      "ACTIVE",
	}
}

func (s paymentTestSuite) TestCreateSuccess() {
	repo := NewPaymentCodeRepository(s.DBConn)
	paymentCodeReq := getMockPaymentCode()

	s.Require().Zero(paymentCodeReq.ID)
	s.Require().Zero(paymentCodeReq.CreatedTime)
	s.Require().Zero(paymentCodeReq.UpdatedTime)
	s.Require().Zero(paymentCodeReq.ExpirationDate)

	err := repo.Create(context.TODO(), paymentCodeReq)
	s.Require().NoError(err)

	expectedPaymentCode := getMockPaymentCode()

	s.Require().Equal(expectedPaymentCode.Name, paymentCodeReq.Name)
	s.Require().Equal(expectedPaymentCode.Status, paymentCodeReq.Status)
	s.Require().Equal(expectedPaymentCode.PaymentCode, paymentCodeReq.PaymentCode)
	s.Require().Equal(paymentCodeReq.CreatedTime.AddDate(50, 0, 0), paymentCodeReq.ExpirationDate)
}

func (s paymentTestSuite) TestCreateFail() {
	repo := NewPaymentCodeRepository(s.DBConn)
	paymentCodeReq := getMockPaymentCode()
	paymentCodeReq.Status = "invalid"

	s.Require().Zero(paymentCodeReq.ID)
	s.Require().Zero(paymentCodeReq.CreatedTime)
	s.Require().Zero(paymentCodeReq.UpdatedTime)
	s.Require().Zero(paymentCodeReq.ExpirationDate)

	// Invalid status
	err := repo.Create(context.TODO(), paymentCodeReq)
	s.Require().Error(err)

	// No name
	newPaymentCodeReq := getInvalidMockPaymentCode()
	err = repo.Create(context.TODO(), newPaymentCodeReq)
	s.Require().Error(err)
}

func (s paymentTestSuite) TestGetSuccess() {
	repo := NewPaymentCodeRepository(s.DBConn)
	paymentCodeReq := getMockPaymentCode()

	s.Require().Zero(paymentCodeReq.ID)
	s.Require().Zero(paymentCodeReq.CreatedTime)
	s.Require().Zero(paymentCodeReq.UpdatedTime)
	s.Require().Zero(paymentCodeReq.ExpirationDate)

	err := repo.Create(context.TODO(), paymentCodeReq)
	s.Require().NoError(err)

	savedPaymentCodeReq, err := repo.GetByID(context.TODO(), paymentCodeReq.ID.String())
	s.Require().NoError(err)

	s.Require().Equal(savedPaymentCodeReq.Name, paymentCodeReq.Name)
	s.Require().Equal(savedPaymentCodeReq.Status, paymentCodeReq.Status)
	s.Require().Equal(savedPaymentCodeReq.PaymentCode, paymentCodeReq.PaymentCode)
	s.Require().Equal(savedPaymentCodeReq.CreatedTime.AddDate(50, 0, 0), paymentCodeReq.ExpirationDate)
}

func (s paymentTestSuite) TestGetFail() {
	repo := NewPaymentCodeRepository(s.DBConn)

	savedPaymentCodeReq, err := repo.GetByID(context.TODO(), "mock-payment-id")
	s.Require().Error(err)
	s.Require().Empty(savedPaymentCodeReq)
}
