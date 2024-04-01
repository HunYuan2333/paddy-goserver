package MeteorologicalData

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type TemperatureCheckData struct {
	location string
	start    time.Time
	end      time.Time
}

func TemperatureCheck(c *gin.Context) {
	var json TemperatureCheckData
	if err := c.ShouldBindJSON(&json); err != nil {
		// 如果解析失败，返回状态码400（Bad Request）和错误信息
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

}
