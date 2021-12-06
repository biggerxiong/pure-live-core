package router

import (
	"github.com/gin-gonic/gin"
	"github.com/iyear/pure-live/api/v1"
	"github.com/iyear/pure-live/conf"
	"github.com/iyear/pure-live/middleware"
	"github.com/iyear/pure-live/util"
)

var r *gin.Engine

func Init() *gin.Engine {
	gin.SetMode(util.IF(conf.C.Server.Debug, gin.DebugMode, gin.ReleaseMode).(string))
	r = gin.New()

	r.Use(middleware.Recovery())
	r.Use(middleware.CORS())
	r.Use(middleware.Static())

	g := r.Group("/api")
	apiV1 := g.Group("/v1")
	{
		apiV1.GET("/live/serve", v1.Serve)
		apiV1.GET("/live/play", v1.Play)
		apiV1.GET("/live/room_info", v1.GetRoomInfo)
		apiV1.GET("/live/play_url", v1.GetPlayURL)
		apiV1.POST("/live/danmaku/send", v1.SendDanmaku)

		apiV1.POST("/fav/list/add", v1.AddFavList)
		apiV1.GET("/fav/list/get_all", v1.GetAllFavLists)
		apiV1.POST("/fav/list/del", v1.DelFavList)
		apiV1.POST("/fav/list/edit", v1.EditFavList)
		apiV1.GET("/fav/list/get", v1.GetFavList)

		apiV1.GET("/fav/get", v1.GetFav)
		apiV1.POST("/fav/add", v1.AddFav)
		apiV1.POST("/fav/del", v1.DelFav)
		apiV1.POST("/fav/edit", v1.EditFav)
	}

	return r
}