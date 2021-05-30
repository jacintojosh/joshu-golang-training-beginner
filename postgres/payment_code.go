package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	errors "joshua-golang-training-beginner/errors"
	"joshua-golang-training-beginner/payment"

	sq "github.com/Masterminds/squirrel"
)

type paymentCodeRepository struct {
	DB *sql.DB
}

func NewPaymentCodeRepository(db *sql.DB) payment.IPaymentCodeRepository {
	return &paymentCodeRepository{
		DB: db,
	}
}

var (
	paymentCodeTableName = "payment_codes"
	paymentCodeColumns   = []string{
		"id",
		"name",
		"payment_code",
		"status",
		"expiration_date",
		"created",
		"updated",
	}
)

func (r paymentCodeRepository) Create(ctx context.Context, p *payment.PaymentCode) (err error) {
	// TODO: Make handling this more graceful by setting name and paymentcode to pointer string type
	if p.Name == "" || p.PaymentCode == "" {
		err = errors.ValidationError("Payment code or name should not be empty.")
		return
	}
	sqlStr, args, err := sq.
		Insert(paymentCodeTableName).
		Columns("name", "payment_code", "status").
		Values(p.Name, p.PaymentCode, p.Status).
		Suffix(fmt.Sprintf("RETURNING %s", strings.Join(paymentCodeColumns, ","))).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return
	}

	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return
	}
	err = tx.QueryRowContext(ctx, sqlStr, args...).
		Scan(&p.ID, &p.Name, &p.PaymentCode, &p.Status, &p.ExpirationDate, &p.CreatedTime, &p.UpdatedTime)

	if err != nil {
		return
	}

	err = tx.Commit()
	if err != nil {
		return
	}

	return
}

func (r paymentCodeRepository) GetByID(ctx context.Context, id string) (p payment.PaymentCode, err error) {
	query := sq.
		Select(paymentCodeColumns...).
		Where(sq.Eq{"id": id}).
		From(paymentCodeTableName).
		PlaceholderFormat(sq.Dollar)

	err = query.
		RunWith(r.DB).
		QueryRowContext(ctx).
		Scan(&p.ID, &p.Name, &p.PaymentCode, &p.Status, &p.ExpirationDate, &p.CreatedTime, &p.UpdatedTime)

	if err == sql.ErrNoRows {
		err = errors.DataNotFoundError("Payment Code ID Not Found")
	}

	return
}
