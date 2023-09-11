package voc

import (
	"github.com/gin-gonic/gin"
	"mock-api/pkg/utils"

	"net/http"
)

func (h *vocHandler) DownloadSentiment(ctx *gin.Context) {
	typeDownload := utils.ConvertStrToInt(ctx.Query("typeDownload"))

	var res interface{}

	if typeDownload == 1 {
		res = "Ok 1"
	} else {
		res = "Ok 2"
	}

	ctx.JSON(http.StatusOK, utils.RestBody{
		Message: "success",
		Data:    res,
	})
	return
}
