package postgresql

import (
	"app/api/models"
	"context"
	"errors"
)



func (r *orderRepo) AddOrderItemNew(ctx context.Context, req *models.CreateOrderItemNew) (int, error) {

	var currentQuantity int
	err := r.db.QueryRow(ctx, "SELECT quantity FROM stocks WHERE product_id = $1", req.ProductID).Scan(&currentQuantity)
	if err != nil {
		return 0,err
	}
	if currentQuantity < req.Quantity {
		return 0, errors.New("not enough quantity in stock")
	}

	query := `
			INSERT INTO order_items(
			order_id, 
			item_id, 
			product_id,
			quantity,
			list_price,
			discount
		)
		VALUES (
			$1, 
			(
				SELECT COALESCE(MAX(item_id), 0) + 1 FROM order_items WHERE order_id = $1
			),
			$2, $3, $4, $5)
	`

	_, err = r.db.Exec(ctx, query,
		req.OrderID,
		req.ProductID,
		req.Quantity,
		req.ListPrice,
		req.Discount,
	)

	if err != nil {
		return 0,err
	}
	_, err = r.db.Exec(ctx, "UPDATE stocks SET quantity = $1 WHERE product_id = $2 and store_id = $3", currentQuantity-req.Quantity, req.ProductID, req.StoreID)
	if err != nil {
		return 0,err
	}



	return 0, err
}
