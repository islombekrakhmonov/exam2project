package postgresql

import (
	"app/api/models"
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"

	_ "github.com/mattn/go-sqlite3"
)


type change struct {
	db *pgxpool.Pool
}

func NewMakeChangeStockRepo(db *pgxpool.Pool) *change {
	return &change{
		db: db,
	}
}

func (r staffRepo) GetSoldProductsByStaff(ctx context.Context, req *models.SoldDate) (resp []models.StaffBySoldItemsOnDay, err error){

	soldDate := req.Date 
	if soldDate == "" {
		soldDate = time.Now().Format("20060102")
	}


	if err != nil {
		return nil, errors.New("error getting")
	}	
	
	var (
		ddd sql.NullString
	)

	insertQuery := `
	SELECT  
	o.order_date, 
	s.first_name,
	s.last_name,  
	p.product_name,
	c.category_name,
	SUM(oi.quantity) AS total_quantity, 
	SUM(oi.list_price * oi.quantity * (1 - oi.discount)) AS total_price 
FROM 
	orders o 
	JOIN order_items oi ON o.order_id = oi.order_id 
	JOIN products p ON oi.product_id = p.product_id 
	JOIN staffs s ON o.staff_id = s.staff_id 
	JOIN categories c ON c.category_id = p.category_id
WHERE 
	o.order_date = $1
GROUP BY 
	o.order_date,
	s.first_name, 
	s.last_name,
	p.product_name,
	c.category_name;
	`


	
	rows, err := r.db.Query(ctx, insertQuery, soldDate)
	if err != nil {
		return resp, err
	}
	defer rows.Close()

	for rows.Next() {
		var staff models.StaffBySoldItemsOnDay
		staff.StaffData = &models.Staff{}
		staff.CategoryData = &models.Category{}
		staff.ProductData = &models.Product{}

		err = rows.Scan(
			&ddd,
			&staff.StaffData.FirstName,
			&staff.StaffData.LastName,
			&staff.ProductData.ProductName,
			&staff.CategoryData.CategoryName,
			&staff.Quantity,
			&staff.TotalPrice,
		)

		if err != nil {
			return resp, err
		}
		resp = append(resp, staff)
	}
	return resp, nil

}
