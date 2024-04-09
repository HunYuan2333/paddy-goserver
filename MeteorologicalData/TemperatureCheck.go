package MeteorologicalData

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"net/http"
	"paddy-goserver/DataBaseConnection"
	"time"
)

var database *sqlx.DB

func init() {
	if err := DataBaseConnection.SetupDatabase(); err != nil {
		panic(err)
	}
	database = DataBaseConnection.GetDatabase()
}

type TemperatureCheckData struct {
	Type     string
	Location string
	Start    Time `json:"start"`
}
type Time struct {
	time.Time
}

func (t *Time) UnmarshalJSON(data []byte) error {
	var timestamp string
	if err := json.Unmarshal(data, &timestamp); err != nil {
		return err
	}
	parsedTime, err := time.Parse(time.RFC3339, timestamp)
	if err != nil {
		return err
	}
	t.Time = parsedTime
	return nil
}
func TemperatureCheck(c *gin.Context) {
	var json TemperatureCheckData
	if err := c.ShouldBindJSON(&json); err != nil {
		// 如果解析失败，返回状态码400（Bad Request）和错误信息
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if json.Start.Time.After(time.Now()) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "start time must be before current time"})
		return
	}
	minimumAllowedTime := time.Date(1981, time.January, 1, 1, 0, 0, 0, time.UTC)
	if json.Start.Time.Before(minimumAllowedTime) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "start time must be on or after 1981-01-01 01:00:00"})
		return
	}

	if json.Type == "year" {
		YearTemperatureCheck(c, json)
		return
	}
}
func YearTemperatureCheck(c *gin.Context, json TemperatureCheckData) {
	start := json.Start.Time
	end := start.AddDate(1, 0, 0)

	// 构建 SQL 查询语句
	queryStr := `
		SELECT DATE_FORMAT(time_month, '%Y-%m') AS month, AVG(temperature) AS avg_monthtem
		FROM MonthAverageTemperature
		WHERE time_month BETWEEN ? AND ?
		GROUP BY month
		ORDER BY month;
	`

	// 执行查询
	rows, err := database.Query(queryStr, start, end)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	temp := make([]float64, 12)
	for rows.Next() {
		var month string
		var avg_monthtem float64
		if err := rows.Scan(&month, &avg_monthtem); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		mon, _ := time.Parse("2006-01", month)
		temp[mon.Month()-1] = avg_monthtem
		if err := rows.Err(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"yData": temp,
	})
}
