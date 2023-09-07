package voc

import (
	"github.com/gin-gonic/gin"
	"mock-api/domain"
	"mock-api/pkg/utils"

	"net/http"
)

func (h *vocHandler) GetTNPS(ctx *gin.Context) {

	res := domain.TNPS{
		OverallTNPSFirstPeriod:  domain.TNPSScore{},
		OverallTNPSSecondPeriod: domain.TNPSScore{},
		TNPSChannel:             domain.TNPSScore{},
		TNPSJourney:             domain.TNPSScore{},
		TNPSTrends:              domain.TNPSScore{},
		TUPCTPChannel:           domain.TNPSScore{},
		TUPDigitalJourney:       domain.TNPSScore{},
	}

	ctx.JSON(http.StatusOK, utils.RestBody{
		Message: "success",
		Data:    res,
	})
	return
}
