package voc

import (
	"github.com/gin-gonic/gin"
	"mock-api/domain"
	"mock-api/pkg/utils"
	"net/http"
)

func (h *vocHandler) GetListTopic(ctx *gin.Context) {

	res := []domain.Topic{
		{
			Id:   1,
			Name: "Pricing",
		},
		{
			Id:   2,
			Name: "Network",
		},
		{
			Id:   3,
			Name: "Performance",
		},
		{
			Id:   4,
			Name: "Product",
		},
	}

	ctx.JSON(http.StatusOK, utils.RestBody{
		Message: "success",
		Data:    res,
	})
	return
}
