package voc

import (
	"github.com/gin-gonic/gin"
	"mock-api/domain"
	"mock-api/pkg/utils"
	"net/http"
)

func (h *vocHandler) GetListArea(ctx *gin.Context) {
	res := []domain.Area{
		{
			Id:   1,
			Name: "Area 1",
		},
		{
			Id:   2,
			Name: "Area 2",
		},
		{
			Id:   3,
			Name: "Area 3",
		},
		{
			Id:   4,
			Name: "Area 4",
		}}

	ctx.JSON(http.StatusOK, utils.RestBody{
		Message: "success",
		Data:    res,
	})
	return
}
