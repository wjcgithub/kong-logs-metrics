package router

import (
	"fmt"
	"github.com/DevWithLin/kong-logs-metrics/config"
	"github.com/DevWithLin/kong-logs-metrics/controller/elastic"
	"github.com/DevWithLin/kong-logs-metrics/controller/login"
	"github.com/DevWithLin/kong-logs-metrics/controller/showlog"

	"github.com/gin-gonic/gin"
)

// Route ?????api??
func Route(router *gin.Engine) {
	apiPrefix := config.Conf.GoConf.APIPrefix

	fmt.Println(apiPrefix)

	api := router.Group(apiPrefix)
	{
		api.POST("/findaggmetrics", agg.FindAggMetrics)
		api.POST("/piechart", agg.PieChar)
		api.POST("/test/queryurlname", agg.QueryURLName)
		api.POST("/checklogin", login.PostCheckLogin)
		api.POST("/showlogs", showlog.ShowLogs)
		api.POST("/findlogdetailbyid", showlog.FindLogDetailByID)
		api.POST("/findlogsbyapiname", showlog.FindLogByAPINameAndDate)
		api.POST("findmatchid", agg.MatchID)
	}
}