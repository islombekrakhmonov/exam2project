package postgresql

import (
	"app/api/models"
	"context"
	"errors"
)


func (r *stockRepo)TransferProduct(ctx context.Context, req *models.MakeATransfer) (int64, error) {

	fromStoreID := req.FromStoreID
	toStoreID := req.ToStoreID
	productID := req.ProductId
	quantity := req.Quantity


	tx, err := r.db.Begin(ctx)
	if err != nil {
		return 0, err
	}

	var fromQuantity int
	ctx = context.Background()

	err = tx.QueryRow(ctx, "SELECT quantity FROM stocks WHERE store_id = $1 AND product_id = $2 FOR UPDATE", fromStoreID, productID).Scan(&fromQuantity)
	if err != nil {
		tx.Rollback(ctx)
		return 0, err
	}

	if fromQuantity < quantity {
		tx.Rollback(ctx)
		return 0, errors.New("not enough quantity in the source store")
	}

	_, err = tx.Exec(ctx, "UPDATE stocks SET quantity = quantity - $1 WHERE store_id = $2 AND product_id = $3", quantity, fromStoreID, productID)
if err != nil {
    tx.Rollback(ctx)
    return 0, err
}

	_, err = tx.Exec(ctx, "UPDATE stocks SET quantity = quantity + $1 WHERE store_id = $2 AND product_id = $3", quantity, toStoreID, productID)
	if err != nil {
		tx.Rollback(ctx)
		return 0,err
	}

	err = tx.Commit(ctx)
	if err != nil {
		tx.Rollback(ctx)
		return 0, err
	}

	return 0, nil
}