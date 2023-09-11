package voc

import (
	"github.com/gin-gonic/gin"
	"mock-api/domain"
	"mock-api/pkg/utils"
	"net/http"
)

func (h *vocHandler) CreateCustomTopic(ctx *gin.Context) {

	var req domain.CreateCustomTopic
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.RestBody{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, utils.RestBody{
		Message: "success",
	})
	return
}
