package handler

import (
	"app/api/models"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)





func (h *Handler) CreateDiscount(c *gin.Context) {

	var createDiscount models.PromoCode

	err := c.ShouldBindJSON(&createDiscount) // parse req body to given type struct
	if err != nil {
		h.handlerResponse(c, "create promo code", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storages.PromoCode().Create(context.Background(), createDiscount)
	if err != nil {
		h.handlerResponse(c, "storage.promocode.create", http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println(id)

	//idN := string(id)

	// resp, err := h.storages.Discount().GetByID(context.Background(), &models.PromoCode{Id: idN})
	// if err != nil {
	// 	h.handlerResponse(c, "storage.discount.getByID", http.StatusInternalServerError, err.Error())
	// 	return
	// }

	h.handlerResponse(c, "create discount", http.StatusCreated, err)
}

func (h *Handler) GetByIdDiscount(c *gin.Context) {

	id := c.Param("id")


	resp, err := h.storages.PromoCode().GetByID(context.Background(), &models.PromoCode{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.brand.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get brand by id", http.StatusCreated, resp)
}

func (h *Handler) GetListDiscount(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list discounts", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list discount", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.storages.PromoCode().GetList(context.Background(), &models.GetListPromoCodetRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.discount.getlist", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list discount response", http.StatusOK, resp)
}

func (h *Handler) DeleteDiscount(c *gin.Context) {

	id := c.Param("id")


	rowsAffected, err := h.storages.PromoCode().Delete(context.Background(), &models.PromoCode{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.discount.delete", http.StatusInternalServerError, err.Error())
		return
	}
	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.discount.delete", http.StatusBadRequest, "now rows affected")
		return
	}

	h.handlerResponse(c, "delete discount", http.StatusNoContent, nil)
}


func (h *Handler) UpdateDiscount(c *gin.Context) {

	var updateDiscount models.PromoCode

	id := c.Param("id")

	err := c.ShouldBindJSON(&updateDiscount)
	if err != nil {
		h.handlerResponse(c, "update discount", http.StatusBadRequest, err.Error())
		return
	}


	updateDiscount.Id = id

	rowsAffected, err := h.storages.PromoCode().Update(context.Background(), &updateDiscount)
	if err != nil {
		h.handlerResponse(c, "storage.discount.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.discount.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.storages.PromoCode().GetByID(context.Background(), &models.PromoCode{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.discount.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "update discount", http.StatusAccepted, resp)
}
