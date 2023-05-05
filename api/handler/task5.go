package handler

import (
	"app/api/models"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)






func (h *Handler)CreateOrderItemNew(c *gin.Context) {
	var createOrderItemNew models.CreateOrderItemNew

	err := c.ShouldBindJSON(&createOrderItemNew) 
	if err != nil {
		h.handlerResponse(c, "create order", http.StatusBadRequest, err.Error())
		return
	}

	_, err = h.storages.Order().AddOrderItemNew(context.Background(), &createOrderItemNew)
	if err != nil {
		h.handlerResponse(c, "storage.order.create", http.StatusInternalServerError, err.Error())
		return
}
h.handlerResponse(c, "create order", http.StatusCreated, "Order Item Added")
}


