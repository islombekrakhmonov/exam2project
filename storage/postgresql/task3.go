package postgresql

import (
	"app/api/models"
	"context"
	"app/pkg/helper"


	"github.com/jackc/pgx/v4/pgxpool"
	"fmt"
)

type promoCodeRepo struct {
	db *pgxpool.Pool
}

func NewPromoCodeRepo(db *pgxpool.Pool) *promoCodeRepo {
	return &promoCodeRepo{
		db: db,
	}
}


func (r *promoCodeRepo) Create(ctx context.Context, req models.PromoCode) (int, error) {

	query := `
		INSERT INTO promo_codes(
			promo_code_id, 
			promo_code_amount,
			promo_code_name, 
			promo_code_type
			
		)
		VALUES (
			$1,
			$2,
			$3,
            $4
		)
		
	`

	_, err := r.db.Exec(ctx ,query,
		req.Id,
		req.Amount,
		req.Name,
	    req.Type,
	)
	

	if err != nil {
		return 0, err
	}

	return 0, nil
}

func (r *promoCodeRepo) GetByID(ctx context.Context, req *models.PromoCode) (*models.PromoCode, error) {

	var (
		query string
		promoCode models.PromoCode
	)

	query = `
		SELECT
			promo_code_id, 
			promo_code_amount,
			promo_code_name,
			promo_code_type
            
		FROM promo_codes
		WHERE promo_code_id = $1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&promoCode.Id,
		&promoCode.Amount,
		&promoCode.Name,
		&promoCode.Type,
	)
	if err != nil {
		return nil, err
	}

	return &promoCode, nil
}


func (r *promoCodeRepo) GetList(ctx context.Context, req *models.GetListPromoCodetRequest) (resp *models.GetListPromoCodeResponse, err error) {

	resp = &models.GetListPromoCodeResponse{}

	var (
		query  string
		filter = " WHERE TRUE "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			promo_code_id,
			promo_code_name,
			promo_code_amount,
            promo_code_type
		FROM promo_codes
	`

	if len(req.Search) > 0 {
		filter += " AND name ILIKE '%' || '" + req.Search + "' || '%' "
	}

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	query += filter + offset + limit

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var promoCode models.PromoCode
		err = rows.Scan(
			&resp.Count,
			&promoCode.Id,
			&promoCode.Name,
			&promoCode.Amount,
            &promoCode.Type,
		)
		if err != nil {
			return nil, err
		}

		resp.PromoCodes = append(resp.PromoCodes, &promoCode)
	}

	return resp, nil
}

func (r *promoCodeRepo) Update(ctx context.Context, req *models.PromoCode) (int64, error) {
	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
		promo_codes
		SET
			promo_code_id = :promo_code_id, 
			promo_code_name = :promo_code_name,
			promo_code_amount = :promo_code_amount,
            promo_code_type = :promo_code_type
		WHERE promo_code_id = :promo_code_id
	`

	params = map[string]interface{}{
		"promo_code_id": req.Id,
        "promo_code_name": req.Name,
        "promo_code_amount": req.Amount,
        "promo_code_type": req.Type,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}


func (r *promoCodeRepo) Delete(ctx context.Context, req *models.PromoCode) (int64, error) {
	query := `
		DELETE 
		FROM promo_codes
		WHERE promo_code_id = $1
	`

	result, err := r.db.Exec(ctx, query, req.Id)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}
