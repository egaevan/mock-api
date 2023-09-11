package voc

import (
	"github.com/gin-gonic/gin"
	"mock-api/domain"
	"mock-api/pkg/utils"
	"net/http"
)

func (h *vocHandler) GetTopicTrend(ctx *gin.Context) {

	res := []domain.Sentiment{
		{
			Description:   "Gila ini sih lemot parah @telkomsel udah 2 harian ini gak bisa donlot",
			ChannelName:   "TNPS",
			Author:        "@wirosablengs",
			Msisdn:        "081111111",
			Location:      "Indonesia",
			Sentiment:     "Negative",
			ARPU:          500000,
			Los:           "2 tahun 2 bulan",
			TelcoBehavior: "Gamer",
		},
	}

	ctx.JSON(http.StatusOK, utils.RestBody{
		Message: "success",
		Data:    res,
	})
	return
}
