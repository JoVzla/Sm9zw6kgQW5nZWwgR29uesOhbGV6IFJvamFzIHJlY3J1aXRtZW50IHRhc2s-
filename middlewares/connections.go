package middlewares

import (
	"gogoapps/models/config"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
)

func InitConnections(cach *cache.Cache, configuration config.Configuration) gin.HandlerFunc {

	return func(c *gin.Context) {
		c.Set("cache", cach)
		c.Set("config", configuration)
		c.Next()
	}

}
