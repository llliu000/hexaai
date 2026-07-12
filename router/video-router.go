package router

import (
	"github.com/QuantumNous/new-api/controller"
	"github.com/QuantumNous/new-api/middleware"

	"github.com/gin-gonic/gin"
)

func SetVideoRouter(router *gin.Engine) {
	// Video proxy: accepts either session auth (dashboard) or token auth (API clients)
	videoProxyRouter := router.Group("/v1")
	videoProxyRouter.Use(middleware.RouteTag("relay"))
	videoProxyRouter.Use(middleware.TokenOrUserAuth())
	{
		videoProxyRouter.GET("/videos/:task_id/content", controller.VideoProxy)
	}

	videoV1Router := router.Group("/v1")
	videoV1Router.Use(middleware.RouteTag("relay"))
	videoV1Router.Use(middleware.TokenAuth(), middleware.Distribute())
	{
		videoV1Router.POST("/video/generations", controller.RelayTask)
		videoV1Router.GET("/video/generations/:task_id", controller.RelayTaskFetch)
		videoV1Router.POST("/videos/:video_id/remix", controller.RelayTask)
	}
	// openai compatible API video routes
	// docs: https://platform.openai.com/docs/api-reference/videos/create
	{
		videoV1Router.POST("/videos", controller.RelayTask)
		videoV1Router.GET("/videos/:task_id", controller.RelayTaskFetch)
	}

	klingV1Router := router.Group("/kling/v1")
	klingV1Router.Use(middleware.RouteTag("relay"))
	klingV1Router.Use(middleware.KlingRequestConvert(), middleware.TokenAuth(), middleware.Distribute())
	{
		klingV1Router.POST("/videos/text2video", controller.RelayTask)
		klingV1Router.POST("/videos/image2video", controller.RelayTask)
		klingV1Router.GET("/videos/text2video/:task_id", controller.RelayTaskFetch)
		klingV1Router.GET("/videos/image2video/:task_id", controller.RelayTaskFetch)
	}

	// Jimeng official API routes - direct mapping to official API format
	jimengOfficialGroup := router.Group("jimeng")
	jimengOfficialGroup.Use(middleware.RouteTag("relay"))
	jimengOfficialGroup.Use(middleware.JimengRequestConvert(), middleware.TokenAuth(), middleware.Distribute())
	{
		// Maps to: /?Action=CVSync2AsyncSubmitTask&Version=2022-08-31 and /?Action=CVSync2AsyncGetResult&Version=2022-08-31
		jimengOfficialGroup.POST("/", controller.RelayTask)
	}

	// doubao https://www.volcengine.com/docs/82379/1521309?lang=zh
	assetV3Router := router.Group("/api/v3/asset")
	assetV3Router.Use(middleware.RouteTag("relay"))
	assetV3Router.Use(middleware.TokenAuth())
	{
		assetController := &controller.AssetController{}
		assetV3Router.POST("", assetController.Action)
		assetV3Router.POST("/", assetController.Action)
		assetV3Router.POST("/manual", assetController.ManualAsset)
	}

	doubaoV3Router := router.Group("/api/v3")
	doubaoV3Router.Use(middleware.RouteTag("relay"))
	doubaoV3Router.Use(middleware.DoubaoRequestConvert(), middleware.TokenAuth(), middleware.Distribute())
	{
		doubaoV3Router.POST("/contents/generations/tasks", controller.RelayTask)
		doubaoV3Router.GET("/contents/generations/tasks/:task_id", controller.RelayTaskFetch)
	}

	// ali wan 2.x https://bailian.console.aliyun.com/cn-beijing/?spm=5176.29597918.J_SEsSjsNv72yRuRFS2VknO.2.7bc6133cyyPFMY&tab=api#/api/?type=model&url=2867393
	wanV1Router := router.Group("/api/v1")
	wanV1Router.Use(middleware.RouteTag("relay"))
	wanV1Router.Use(middleware.WanRequestConvert(), middleware.TokenAuth(), middleware.Distribute())
	{
		wanV1Router.POST("/services/aigc/video-generation/video-synthesis", controller.RelayTask)
		wanV1Router.GET("/tasks/:task_id", controller.RelayTaskFetch)
	}
}
