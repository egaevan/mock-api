package voc

import (
	"github.com/gin-gonic/gin"
	"mock-api/domain"
	"mock-api/pkg/utils"
	"net/http"
)

func (h *vocHandler) GetListCustomerType(ctx *gin.Context) {
	res := []domain.CustomerType{
		{
			Id:       1,
			TypeName: "High Revenue Contributor (HRC)",
		},
		{
			Id:       2,
			TypeName: "P50",
		},
		{
			Id:       3,
			TypeName: "P75",
		},
		{
			Id:       4,
			TypeName: "P95",
		},
		{
			Id:       5,
			TypeName: "The rest",
		},
		{
			Id:       6,
			TypeName: "Blank",
		},
		{
			Id:       7,
			TypeName: "High Value Customer (HVC)",
		},
		{
			Id:       8,
			TypeName: "Silver",
		},
		{
			Id:       8,
			TypeName: "Gold",
		},
		{
			Id:       8,
			TypeName: "Diamond",
		},
		{
			Id:       8,
			TypeName: "Platinum",
		}}

	ctx.JSON(http.StatusOK, utils.RestBody{
		Message: "success",
		Data:    res,
	})
	return
}
