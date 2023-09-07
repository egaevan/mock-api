package voc

import (
	"github.com/gin-gonic/gin"
	"mock-api/domain"
	"mock-api/pkg/utils"
	"net/http"
)

func (h *vocHandler) GetListChannel(ctx *gin.Context) {
	res := []domain.Channel{
		{
			Id:          1,
			ChannelName: "Grapari",
		},
		{
			Id:          2,
			ChannelName: "MyGrapari",
		},
		{
			Id:          3,
			ChannelName: "Call Center",
		},
		{
			Id:          4,
			ChannelName: "Network (OTA)",
		},
		{
			Id:          5,
			ChannelName: "MyTelkomsel",
		},
		{
			Id:          6,
			ChannelName: "Prepaid Product",
		},
		{
			Id:          7,
			ChannelName: "MAXStream",
		},
		{
			Id:          8,
			ChannelName: "Veronika",
		}}

	ctx.JSON(http.StatusOK, utils.RestBody{
		Message: "success",
		Data:    res,
	})
	return
}
