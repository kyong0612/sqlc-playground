package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/kyong0612/sqlc-playground/persistence"
)

func IncrementUserAges(ctx context.Context, conn *pgx.Conn, q *persistence.Queries, id int32) error {
	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}
	qWithTx := q.WithTx(tx)
	u, err := qWithTx.GetUser(ctx, id)
	if err != nil {
		return err
	}
	err = qWithTx.UpdateUserAges(ctx, persistence.UpdateUserAgesParams{
		ID:  u.ID,
		Age: u.Age + 1,
	})
	if err != nil {
		return err
	}
	if err := tx.Commit(ctx); err != nil {
		return err
	}
	return nil
}
