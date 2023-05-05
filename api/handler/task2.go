package handler

import (
	"app/api/models"
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)


func (h *Handler) GetSold(c *gin.Context) {
	var err error

	date := c.Param("date")

	idInt := string(date)
	if err != nil {
		h.handlerResponse(c, "storage.staff.getSoldProducts ", http.StatusBadRequest, "date is not incorrect")
		return
	}

	resp, err := h.storages.Staff().GetSoldProductsByStaff(context.Background(), &models.SoldDate{Date: idInt})
	if err != nil {
		h.handlerResponse(c, "storage.staff.getSoldProducts", http.StatusInternalServerError, errors.New("no products were sold on that day"))
		return
	}
	
	h.handlerResponse(c, "get sold products by staff", http.StatusCreated, resp)
}