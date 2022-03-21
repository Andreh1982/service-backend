package routes

import (
	"service-backend/controllers"
	"service-backend/database"
	"service-backend/metric"
	"service-backend/middleware"

	"service-backend/shared"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func HandleRequest() {

	router := gin.Default()
	gin.SetMode("release")

	shared.LogCustom([]string{"Inicializando métricas"}, "info")

	metricService, _ := metric.NewPrometheusService()
	router.Use(middleware.Metrics(metricService))
	router.GET("/metrics", func(c *gin.Context) {
		promHandler := promhttp.Handler()
		promHandler.ServeHTTP(c.Writer, c.Request)
	})

	shared.LogCustom([]string{"Criando rotas"}, "info")

	api := &controllers.APIEnv{
		DB: database.GetDB(),
	}

	router.GET("/buyers", api.GetBuyers)
	router.GET("/buyer/:id", api.GetBuyer)
	router.POST("addbuyer", api.CreateBuyer)
	router.PUT("/updatebuyer/:id", api.UpdateBuyer)
	router.DELETE("/delbuyer/:id", api.DeleteBuyer)

	router.GET("/sellers", api.GetSellers)
	router.GET("/seller/:id", api.GetSeller)
	router.POST("/addseller", api.CreateSeller)
	router.PUT("/updateseller/:id", api.UpdateSeller)
	router.DELETE("/delseller/:id", api.DeleteSeller)

	shared.LogCustom([]string{"Iniciando a API"}, "info")

	router.Run(":9990")

}
