package routes

import (
	"gogoapps/controllers/weather"

	"github.com/gin-gonic/gin"
)

//Weather routes
func Weather(r *gin.Engine) {
	r.GET("/cities", weather.GetWeather)
}
