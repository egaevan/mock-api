package voc

import (
	"github.com/gin-gonic/gin"
	"mock-api/domain"
	"mock-api/pkg/utils"
	"net/http"
)

func (h *vocHandler) GetGeoMap(ctx *gin.Context) {

	res := domain.GeoMapPage{
		GeoMap: []domain.GeoMap{
			{
				TNPSScore: domain.TNPSScore{
					TotalRespondent: 0,
					Promotor:        0,
					Passive:         0,
					Detractor:       0,
				},
				Location: domain.Location{
					Id:           0,
					LocationName: "",
				},
			},
		},
		GeoByArea: domain.GeoByArea{
			Area: domain.Area{
				Id:       0,
				Location: nil,
			},
			TopicRatio: []domain.TopicRatio{
				{
					Topic: domain.Topic{
						Id:   0,
						Name: "",
					},
					Ratio: 0,
				},
			},
		},
	}

	ctx.JSON(http.StatusOK, utils.RestBody{
		Message: "success",
		Data:    res,
	})
	return
}
