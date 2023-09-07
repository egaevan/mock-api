package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	dugemHttp "mock-api/delivery/dugem"
	vocHttp "mock-api/delivery/voc"
	"net/http"
)

var (
	app      *gin.Engine
	basePath = "/api/mock"
	vocHD    = vocHttp.NewHandler()
	dugemHD  = dugemHttp.NewHandler()
)

func init() {
	app = gin.Default()
	app.NoRoute(ErrRouter)
	route := app.Group(basePath)

	route.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	dugemRoute := route.Group("/dugem")
	{
		dugemRoute.GET("/pdf", dugemHD.GetPdf)
	}

	vocRoute := route.Group("/voc")
	{
		tnps := vocRoute.Group("/tnps")
		{
			tnps.GET("/", vocHD.GetTNPS)
			tnps.GET("/metric", vocHD.GetTNPSMetric)
			tnps.GET("/metric/download", vocHD.DownloadMetric)
		}

		sentiment := vocRoute.Group("/sentiment")
		{
			sentiment.GET("/", vocHD.GetSentiment)
			sentiment.GET("/timeline", vocHD.GetSentimentTimeline)
			sentiment.GET("/detail-profile", vocHD.GetSentimentProfile)
			sentiment.GET("/detail-profile/download", vocHD.DownloadInteraction)
		}

		list := vocRoute.Group("/list")
		{
			list.GET("/channel", vocHD.GetListChannel)
			list.GET("/customer-type", vocHD.GetListCustomerType)
			list.GET("/journey", vocHD.GetListJourney)
			list.GET("/area/location", vocHD.GetListLocation)
		}
	}
}

func ErrRouter(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"errors": "this page could not be found",
	})
}

// Entrypoint
func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
