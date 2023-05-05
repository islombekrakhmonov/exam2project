package handler

import (
	"app/api/models"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)



func (h *Handler)TransferProductHandler(r *gin.Context) {

	var req models.MakeATransfer
	if err := r.ShouldBindJSON(&req); err != nil {
        r.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

ctx := context.Background()
	if _, err := h.storages.Stock().TransferProduct(ctx, &req); err != nil {
		r.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	r.JSON(http.StatusOK, gin.H{"message": "Product transferred successfully"})
}
