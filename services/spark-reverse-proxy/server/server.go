package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"reverse-proxy/handlers"
	"time"
)

func Run(config Config) {
	port := config.Port

	router := gin.Default()

	apiConfig := handlers.ApiConfig{
		SparkApplicationNamespace: config.SparkApplicationNamespace,
		SparkUIServiceUrl:         config.SparkUIServiceUrl,
		ModifyRedirectUrl:         config.ModifyRedirectUrl,
	}

	router.GET("/version", handlers.Version)

	now := time.Now()
	router.GET("/health",
		func(context *gin.Context) {
			handlers.HealthCheck(context, now.Format(time.RFC3339))
		})

	router.GET(fmt.Sprintf("/%s/*path", config.ProxyBaseUri),
		func(context *gin.Context) {
			handlers.ServeSparkUI(context, &apiConfig, fmt.Sprintf("/%s", config.ProxyBaseUri))
		})

	router.Run(fmt.Sprintf(":%d", port))
}
