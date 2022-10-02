package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type PaymentRepository struct {
	db *sqlx.DB
}

func NewPaymentRepository(db *sqlx.DB) *PaymentRepository {
	return &PaymentRepository{
		db: db,
	}
}

func (r *PaymentRepository) CreatePaymentAccount(accountId int, stripeId string) error {
	query := fmt.Sprintf(`UPDATE %s SET stripe_id=$1 WHERE id =$2`, artistTable)
	_, err := r.db.Exec(query, stripeId, accountId)

	return err
}
