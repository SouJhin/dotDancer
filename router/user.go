package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitUserRoutes() {
	RegistRoute(func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup) {
		rgPublic.POST("/login", func(context *gin.Context) {
			context.AbortWithStatusJSON(http.StatusOK, gin.H{
				"msg": "user login",
			})
		})
		rgAuthUser := rgAuth.Group("user")
		rgAuthUser.GET("", func(context *gin.Context) {
			context.AbortWithStatusJSON(http.StatusOK, gin.H{
				"data": []map[string]interface{}{
					{
						"id":   1,
						"name": "bob",
					},
					{
						"id":   2,
						"name": "fuck u",
					},
				},
			})
		})
		rgAuthUser.GET("/id", func(context *gin.Context) {
			context.AbortWithStatusJSON(http.StatusOK, gin.H{
				"msg": "fuck u",
			})
		})
	})
}
