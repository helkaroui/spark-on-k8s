package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	handlers "reverse-proxy/pkg/handlers"
)

func Run(config Config) {
	port := config.Port

	router := gin.Default()

	apiConfig := handlers.ApiConfig{
		SparkApplicationNamespace: config.SparkApplicationNamespace,
		SparkUIServiceUrl:         config.SparkUIServiceUrl,
		ModifyRedirectUrl:         config.ModifyRedirectUrl,
	}

	router.GET("/health", handlers.HealthCheck)

	router.GET("/sparkui/*path",
		func(context *gin.Context) {
			handlers.ServeSparkUI(context, &apiConfig, "/sparkui")
		})

	router.Run(fmt.Sprintf(":%d", port))
}
