package voc

import (
	"github.com/gin-gonic/gin"
	"mock-api/domain"
	"mock-api/pkg/utils"
	"net/http"
)

func (h *vocHandler) GetWorldCloud(ctx *gin.Context) {

	res := []domain.Word{
		{
			Name:  "Telkomsel",
			Total: 1000,
		},
		{
			Name:  "Jaringan",
			Total: 500,
		},
		{
			Name:  "Sinyal",
			Total: 100,
		},
	}

	ctx.JSON(http.StatusOK, utils.RestBody{
		Message: "success",
		Data:    res,
	})
	return
}
