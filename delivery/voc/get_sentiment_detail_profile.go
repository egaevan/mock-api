package voc

import (
	"github.com/gin-gonic/gin"
	"mock-api/domain"
	"mock-api/pkg/utils"
	"net/http"
	"time"
)

func (h *vocHandler) GetSentimentProfile(ctx *gin.Context) {

	res := domain.SentimentDetailProfile{
		UserProfile: []domain.SocialMedia{
			{
				SocialMediaName: "Twitter",
				Username:        "@wirosablengs",
				Author:          "",
				Followers:       24,
				Following:       1923,
				Friends:         0,
			},
		},
		UserData: domain.UserData{
			Msisdn:        "08217242839",
			ARPU:          500000,
			Los:           "Negative",
			TelcoBehavior: "Gamer",
			Location:      "Indonesia",
			NPS:           2,
			CSI:           "",
			TNPS:          1,
			CTPs:          "",
		},
		SentimentProfile: []domain.SentimentProfile{
			{
				Description: "Gila ini sih lemot parah @telkomsel udah 2 harian ini gak bisa donlot",
				ChannelName: "tNPS",
				Date:        time.Time{},
				Author:      "@wirosablengs",
				Location:    "Indonesia",
				Sentiment:   "Negative",
				ThreadUrl:   "twitter.com",
			},
		},
	}

	ctx.JSON(http.StatusOK, utils.RestBody{
		Message: "success",
		Data:    res,
	})
	return
}
