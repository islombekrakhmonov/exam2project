package handler

import (
	"app/api/models"
	"context"
	"net/http"
	"github.com/gin-gonic/gin"
)

func (h *Handler)GetSum(r *gin.Context) {

	var req models.CountSumRequest
	if err := r.ShouldBindJSON(&req); err != nil {
        r.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    ctx := context.Background()
	resp, err := h.storages.Order().GetTheSum(ctx, &req); 
	if err != nil {
		r.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	
	body := gin.H{"total_sum": resp.TotalSum}
    r.JSON(http.StatusOK, gin.H{"message": "Sum calculated successfully", "data": body})
}