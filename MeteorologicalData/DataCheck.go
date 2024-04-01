package MeteorologicalData

import "github.com/gin-gonic/gin"

func DataCheck(c *gin.Context) {
	Type := c.Param("Type")
	if Type == "temperature" {
		TemperatureCheck(c)
	}
}
