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
		tnpsRoute := vocRoute.Group("/tnps")
		{
			tnpsRoute.GET("/", vocHD.GetTNPS)
			tnpsRoute.GET("/metric", vocHD.GetTNPSMetric)
			tnpsRoute.GET("/metric/download", vocHD.DownloadMetric)
		}

		sentimentRoute := vocRoute.Group("/sentiment")
		{
			sentimentRoute.GET("/", vocHD.GetSentiment)
			sentimentRoute.GET("/timeline", vocHD.GetSentimentTimeline)
			sentimentRoute.GET("/download", vocHD.DownloadSentiment)
			sentimentRoute.GET("/detail-profile", vocHD.GetSentimentProfile)
			sentimentRoute.GET("/detail-profile/download", vocHD.DownloadInteraction)
		}

		geoMapRoute := vocRoute.Group("/geo")
		{
			geoMapRoute.GET("/", vocHD.GetGeoMap)
		}

		topicRoute := vocRoute.Group("/topic")
		{
			topicRoute.GET("/detail", vocHD.GetDetailTopic)
			topicRoute.GET("/detail/download", vocHD.DownloadDetailTopic)
			topicRoute.GET("/trends", vocHD.GetTopicTrend)
			topicRoute.GET("/custom", vocHD.GetTopicTrend)
			topicRoute.GET("/custom/download", vocHD.DownloadCustomTopic)
			topicRoute.POST("/custom", vocHD.CreateCustomTopic)
			topicRoute.GET("/analyzed", vocHD.GetTopicAnalyzed)
			topicRoute.GET("/analyzed/download", vocHD.DownloadTopicAnalyzed)
		}

		worldCloudRoute := vocRoute.Group("/world-cloud")
		{
			worldCloudRoute.GET("/", vocHD.GetWorldCloud)
		}

		list := vocRoute.Group("/list")
		{
			list.GET("/customer-type", vocHD.GetListCustomerType)
			list.GET("/area", vocHD.GetListArea)
			list.GET("/channel", vocHD.GetListChannel)
			list.GET("/journey", vocHD.GetListJourney)
			list.GET("/source", vocHD.GetListSource)
			list.GET("/sentiment", vocHD.GetListSentimentCategory)
			list.GET("/area/location", vocHD.GetListLocation)
			list.GET("/topic", vocHD.GetListTopic)
			list.GET("/nps-score", vocHD.GetListTopic)
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
