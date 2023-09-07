package voc

import (
	"github.com/gin-gonic/gin"
	"mock-api/domain"
	"mock-api/pkg/utils"
	"net/http"
)

func (h *vocHandler) GetListJourney(ctx *gin.Context) {

	res := []domain.Journey{
		{
			Id:          1,
			Category:    "Interact Journey",
			JourneyName: "Interact",
		},
		{
			Id:          2,
			Category:    "Interact Journey",
			JourneyName: "Choose",
		},
		{
			Id:          3,
			Category:    "Using Journey",
			JourneyName: "Consume",
		},
		{
			Id:          4,
			Category:    "Using Journey",
			JourneyName: "Get Help",
		},
		{
			Id:          5,
			Category:    "Using Journey",
			JourneyName: "Pay",
		},
		{
			Id:          6,
			Category:    "Sharing Journey",
			JourneyName: "Leave",
		}}

	ctx.JSON(http.StatusOK, utils.RestBody{
		Message: "success",
		Data:    res,
	})
	return
}
