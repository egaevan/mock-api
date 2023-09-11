package voc

import (
	"github.com/gin-gonic/gin"
	"mock-api/domain"
	"mock-api/pkg/utils"
	"net/http"
)

func (h *vocHandler) GetDetailTopic(ctx *gin.Context) {

	res := domain.TopicDetail{
		TotalPhrase: 1573,
		SentimentScore: domain.SentimentScore{
			Total:    1656,
			Positive: 29,
			Neutral:  10,
			Mixed:    0,
			Negative: 1617,
		},
	}

	ctx.JSON(http.StatusOK, utils.RestBody{
		Message: "success",
		Data:    res,
	})
	return
}
