package voc

import (
	"github.com/gin-gonic/gin"
	"mock-api/domain"
	"mock-api/pkg/utils"
	"time"

	"net/http"
)

func (h *vocHandler) GetSentimentProfile(ctx *gin.Context) {

	res := domain.SentimentData{
		SentimentTimeline: []domain.SentimentScoreTimeline{
			{
				Total:    0,
				Positive: 0,
				Neutral:  0,
				Mixed:    0,
				Negative: 0,
				Date:     time.Time{},
			},
		},
		SentimentBreakdown: domain.SentimentScore{
			Total:    0,
			Positive: 0,
			Neutral:  0,
			Mixed:    0,
			Negative: 0,
		},
		SentimentScore: 40,
	}

	ctx.JSON(http.StatusOK, utils.RestBody{
		Message: "success",
		Data:    res,
	})
	return
}
