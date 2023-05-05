package postgresql

import (
	"app/api/models"
	"context"
	"database/sql"
	"fmt"
	"app/pkg/helper"

)


func(r orderRepo) GetTheSum(ctx context.Context, req *models.CountSumRequest) (resp models.CoustSumResponse, err error) {


	discountID := req.DiscountId
	
	
	var discountAmount float64

	if discountID != "" {
		var (
			query  string
			params map[string]interface{}
		)

		query = `
		UPDATE orders SET promo_code_id = $1 WHERE order_id = $2`

		params = map[string]interface{}{
			"promo_code_id":   req.DiscountId,
			"order_id": req.OrderId,
		}
	
		query, args := helper.ReplaceQueryParams(query, params)
	
		_, err := r.db.Exec(ctx, query, args...)
		if err != nil {
			return resp, err
		}
		

		row := r.db.QueryRow(ctx, "SELECT promo_code_amount FROM promo_codes WHERE promo_code_id = $1", discountID)
		err = row.Scan(&discountAmount)
	

		if err == sql.ErrNoRows {
            return resp, nil

        } else if err != nil {
			fmt.Println(err)
		}
		var TotalSum float64


		selectQuery := `SELECT SUM(oi.list_price * oi.quantity * (1 - oi.discount) * (1 - p.promo_code_amount)) as total_sum
		FROM orders o
		JOIN order_items oi ON o.order_id = oi.order_id
		JOIN promo_codes p ON o.promo_code_id = p.promo_code_id
		WHERE o.order_id = $1;`


		err = r.db.QueryRow(ctx, selectQuery, req.OrderId).Scan(
			&TotalSum,
		)
		
		
		if err != nil {
			return resp, err
		}
	
		return models.CoustSumResponse{
			TotalSum: TotalSum,
		}, nil

	} else {
		var TotalSum float64

		selectQuery := `SELECT SUM(oi.list_price * oi.quantity * (1 - oi.discount)) as total_sum
		FROM orders o
		JOIN order_items oi ON o.order_id = oi.order_id
		WHERE o.order_id = $1;`

		err = r.db.QueryRow(ctx, selectQuery, req.OrderId).Scan(
			&TotalSum,
		)
		
		
		if err != nil {
			return resp, err
		}
	
		return models.CoustSumResponse{
			TotalSum: TotalSum,
		}, nil
	}
}

	

