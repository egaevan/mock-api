package voc

import (
	"github.com/gin-gonic/gin"
	"mock-api/domain"
	"mock-api/pkg/utils"
	"net/http"
)

func (h *vocHandler) GetTNPSMetric(ctx *gin.Context) {

	res := []domain.TNPSMetric{
		{
			ChannelName: "Grapari",
			Interact:    1718,
			Choose:      1524,
			Consume:     0,
			GetHelp:     56,
			Pay:         169,
			Leave:       124,
		},
		{
			ChannelName: "MyGrapari",
			Interact:    0,
			Choose:      0,
			Consume:     0,
			GetHelp:     0,
			Pay:         0,
			Leave:       0,
		},
		{
			ChannelName: "Call Center",
			Interact:    0,
			Choose:      0,
			Consume:     0,
			GetHelp:     0,
			Pay:         0,
			Leave:       0,
		},
		{
			ChannelName: "Network (OTA)",
			Interact:    0,
			Choose:      0,
			Consume:     0,
			GetHelp:     0,
			Pay:         0,
			Leave:       0,
		},
		{
			ChannelName: "MyTelkomsel",
			Interact:    0,
			Choose:      0,
			Consume:     0,
			GetHelp:     0,
			Pay:         0,
			Leave:       0,
		},
		{
			ChannelName: "Prepaid Product",
			Interact:    0,
			Choose:      0,
			Consume:     0,
			GetHelp:     0,
			Pay:         0,
			Leave:       0,
		},
		{
			ChannelName: "MAXStream",
			Interact:    0,
			Choose:      0,
			Consume:     0,
			GetHelp:     0,
			Pay:         0,
			Leave:       0,
		},
		{
			ChannelName: "Veronika",
			Interact:    0,
			Choose:      0,
			Consume:     0,
			GetHelp:     0,
			Pay:         0,
			Leave:       0,
		}}

	ctx.JSON(http.StatusOK, utils.RestBody{
		Message: "success",
		Data:    res,
	})
	return
}
