package weather

import (
	"gogoapps/logger"
	"gogoapps/models/config"
	"gogoapps/models/weathermodel"
	"gogoapps/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
)

var cities []string

//GetWeather get weather of cities
func GetWeather(c *gin.Context) {

	//get values from query strings
	cities = strings.Split(c.Query("values"), ",")

	result := []weathermodel.Weather{}
	cach := c.MustGet("cache").(*cache.Cache)
	configuration := c.MustGet("config").(config.Configuration)

	for i := range cities {
		if utils.CheckCache(cities[i], cach) == nil {
			tmp, err := utils.GetCityData(cities[i], configuration)
			if err != nil {
				logger.Log.Error(err)
			} else {
				result = append(result, tmp)
			}
		}
	}

	c.JSON(http.StatusOK, result)
}
