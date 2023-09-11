package voc

import (
	"github.com/gin-gonic/gin"
	"mock-api/domain"
	"mock-api/pkg/utils"
	"net/http"
)

func (h *vocHandler) GetTopicAnalyzed(ctx *gin.Context) {

	res := []domain.TopicData{
		{
			Id:   4,
			Name: "Performance",
			SentimentScore: domain.SentimentScore{
				Total:    100,
				Positive: 10,
				Neutral:  15,
				Mixed:    10,
				Negative: 75,
			},
		},
		{
			Id:   5,
			Name: "Frequency",
			SentimentScore: domain.SentimentScore{
				Total:    150,
				Positive: 60,
				Neutral:  15,
				Mixed:    10,
				Negative: 75,
			},
		},
		{
			Id:   6,
			Name: "Connectivity",
			SentimentScore: domain.SentimentScore{
				Total:    75,
				Positive: 5,
				Neutral:  50,
				Mixed:    10,
				Negative: 10,
			},
		},
	}

	ctx.JSON(http.StatusOK, utils.RestBody{
		Message: "success",
		Data:    res,
	})
	return
}
