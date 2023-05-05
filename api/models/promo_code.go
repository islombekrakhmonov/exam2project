package models

type PromoCode struct{
	Id string `json:"discount_id"`
	Amount float64 `json:"discount_amount"`
	Name string `json:"discount_name"`
	Type string `json:"discount_type"`
} 


type GetListPromoCodetRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListPromoCodeResponse struct {
		Count  int      `json:"count"`
		PromoCodes []*PromoCode `json:"discounts"`
}