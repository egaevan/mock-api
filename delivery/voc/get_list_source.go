package voc

import (
	"github.com/gin-gonic/gin"
	"mock-api/domain"
	"mock-api/pkg/utils"
	"net/http"
)

func (h *vocHandler) GetListSource(ctx *gin.Context) {

	res := []domain.Source{
		{
			Id:          1,
			Name:        "MyTelkomsel",
			Image:       "",
			Description: "Play Store - Android",
		},
		{
			Id:          2,
			Name:        "MyTelkomsel",
			Image:       "",
			Description: "App Store - IOS",
		},
		{
			Id:          3,
			Name:        "Twitter",
			Image:       "",
			Description: "",
		},
		{
			Id:          4,
			Name:        "Instagram",
			Image:       "",
			Description: "",
		},
		{
			Id:          5,
			Name:        "Facebook",
			Image:       "",
			Description: "",
		},
		{
			Id:          6,
			Name:        "Tiktok",
			Image:       "",
			Description: "",
		},
		{
			Id:          7,
			Name:        "tNPS",
			Image:       "",
			Description: "",
		},
	}

	ctx.JSON(http.StatusOK, utils.RestBody{
		Message: "success",
		Data:    res,
	})
	return
}
