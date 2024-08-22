package routes

import (
	"net/http"
	"tts/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(CORSMiddleware())

	// 加载模板文件
	router.LoadHTMLGlob("templates/*")

	router.GET("/voices", handlers.GetVoiceList)
	router.POST("/tts", handlers.SynthesizeVoicePost)
	router.GET("/tts", handlers.SynthesizeVoice)
	router.GET("/", handlers.Index)

	return router
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // 允许所有域名请求
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
