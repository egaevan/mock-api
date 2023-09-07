package voc

import (
	"github.com/gin-gonic/gin"
	"mock-api/domain"
	"mock-api/pkg/utils"
	"net/http"
)

func (h *vocHandler) GetListLocation(ctx *gin.Context) {
	res := []domain.Area{
		{
			Id: 1,
			Location: []domain.Location{
				{
					Id:           1,
					LocationName: "Sumbagut",
				},
				{
					Id:           2,
					LocationName: "Sumbagteng",
				},
				{
					Id:           3,
					LocationName: "Sumbagsel",
				},
			},
		},
		{
			Id: 2,
			Location: []domain.Location{
				{
					Id:           4,
					LocationName: "Western Jabodetabek",
				},
				{
					Id:           5,
					LocationName: "Central Jabodetabek",
				},
				{
					Id:           6,
					LocationName: "East Jabodetabek",
				},
				{
					Id:           7,
					LocationName: "Jawa Barat",
				},
			},
		},
		{
			Id: 3,
			Location: []domain.Location{
				{
					Id:           8,
					LocationName: "Jawa Tengah",
				},
				{
					Id:           9,
					LocationName: "Jawa Timur",
				},
				{
					Id:           10,
					LocationName: "Bali Nusra",
				},
			},
		},
		{
			Id: 4,
			Location: []domain.Location{
				{
					Id:           11,
					LocationName: "Kalimantan",
				},
				{
					Id:           12,
					LocationName: "Sulawesi",
				},
				{
					Id:           13,
					LocationName: "Puma",
				},
			},
		}}

	ctx.JSON(http.StatusOK, utils.RestBody{
		Message: "success",
		Data:    res,
	})
	return
}
