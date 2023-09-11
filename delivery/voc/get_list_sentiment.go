package voc

import (
	"github.com/gin-gonic/gin"
	"mock-api/domain"
	"mock-api/pkg/utils"
	"net/http"
)

func (h *vocHandler) GetListSentimentCategory(ctx *gin.Context) {

	res := []domain.SentimentCategory{
		{
			Id:   1,
			Name: "Positive",
		},
		{
			Id:   2,
			Name: "Neutral",
		},
		{
			Id:   3,
			Name: "Negative",
		},
	}

	ctx.JSON(http.StatusOK, utils.RestBody{
		Message: "success",
		Data:    res,
	})
	return
}
