package controller

import (
	"fmt"
	_ "gateway/docs"
	"gateway/utils"
	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
	"time"
)

// 初始化路由信息
func InitRouter() *gin.Engine {
	conf := utils.GetConf()
	// 设置运行模式
	if conf.Http.RunMode == "dev" {
		gin.SetMode(gin.DebugMode)
	} else if conf.Http.RunMode == "pro" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		fmt.Println("conf.http.run_mode error,please check.")
	}

	r := gin.Default()
	r.Use(cors.Default())

	store := persistence.NewInMemoryStore(time.Second)

	// 定义路由信息
	if gin.IsDebugging() {
		r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{"msg": "ok"})
	})
	api := r.Group("/v1")
	{
		cluster := api.Group("/cluster")
		{
			controller := ClusterInfoController{}
			cluster.POST("/save", controller.NewOrModify)
			cluster.GET("/query/:id", cache.CachePage(store, time.Minute, controller.Query))
			cluster.DELETE("/delete/:id", controller.Delete)
		}

	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": 404, "message": "Not Found"})
	})
	return r
}

// 启动Http服务
func Start() {
	conf := utils.GetConf()
	fmt.Println("start http server on: 0.0.0.0:", conf.Http.ListenPort)
	routers := InitRouter()
	endPoint := fmt.Sprintf(":%d", conf.Http.ListenPort)
	maxHeaderBytes := 1 << 20
	server := &http.Server{
		Addr:           endPoint,
		Handler:        routers,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: maxHeaderBytes,
	}

	server.ListenAndServe()
}
