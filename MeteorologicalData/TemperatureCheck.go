package MeteorologicalData

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"log"
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
	End      Time `json:"end"`
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
	var temperatureCheckData TemperatureCheckData
	if err := c.ShouldBindJSON(&temperatureCheckData); err != nil {
		// 如果解析失败，返回状态码400（Bad Request）和错误信息
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if temperatureCheckData.Start.Time.After(time.Now()) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "start time must be before current time"})
		return
	}
	minimumAllowedTime := time.Date(1981, time.January, 1, 1, 0, 0, 0, time.UTC)
	if temperatureCheckData.Start.Time.Before(minimumAllowedTime) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "start time must be on or after 1981-01-01 01:00:00"})
		return
	}

	switch temperatureCheckData.Type {
	case "year":
		YearTemperatureCheck(c, temperatureCheckData)
	case "month":
		MonthTemperatureCheck(c, temperatureCheckData)
	case "day":
		DayTemperatureCheck(c, temperatureCheckData)
	case "manyyear":
		InterannualTemperatureCheck(c, temperatureCheckData)

	}
}

// DayTemperatureCheck 函数按小时遍历24小时的温度数据
func DayTemperatureCheck(c *gin.Context, json TemperatureCheckData) {
	start := json.Start.Time
	end := start.Add(24 * time.Hour) // 遍历24小时后的时刻

	// 构建 SQL 查询语句
	queryStr := `
		SELECT HOUR(time) AS hour, AVG(temperature) AS avg_hourlytem
		FROM Temperature
		WHERE time BETWEEN ? AND ?
		GROUP BY hour
		ORDER BY hour;
	`
	rows, err := database.Query(queryStr, start, end)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println("Error executing query:", err)
		return
	}
	defer rows.Close()

	// 初始化24小时的温度数组
	temp := make([]float64, 24)

	for rows.Next() {
		var hour int
		var avg_hourlytem float64
		if err := rows.Scan(&hour, &avg_hourlytem); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		temp[hour] = avg_hourlytem
		if err := rows.Err(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"yData": temp, // 使用 "hData" 以区别于年数据，表示小时数据
	})
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
func MonthTemperatureCheck(c *gin.Context, json TemperatureCheckData) {
	start := json.Start.Time
	end := start.AddDate(0, 1, 0)

	// 使用临时解决方案，将 time_day 显式转换为字符串
	queryStr := `
		SELECT DATE_FORMAT(time_day, '%Y-%m-%d') AS time_day_str, temperature
		FROM Paddy.DayAverageTemperature
		WHERE time_day BETWEEN ? AND ? AND DAY(time_day) MOD 5 = 0
		ORDER BY time_day
	`

	// 执行查询
	rows, err := database.Query(queryStr, start, end)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	temps := []float64{}
	for rows.Next() {
		var timeStr string
		var temp float64

		if err := rows.Scan(&timeStr, &temp); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		_, err := time.Parse("2006-01-02", timeStr)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		temps = append(temps, temp)
		log.Println(temp)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if len(temps) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"yData": temps,
		})
		return
	}
	/*	log.Println(temps)*/
	c.JSON(http.StatusOK, gin.H{
		"yData": temps,
	})
}
func InterannualTemperatureCheck(c *gin.Context, json TemperatureCheckData) {
	start := json.Start.Time
	end := json.End.Time
	start = start.AddDate(1, 0, 0)
	end = end.AddDate(1, 0, 0)
	log.Println(start)
	log.Println(end)
	// 构建 SQL 查询语句
	queryStr := `
		SELECT temperature
		FROM YearAverageTemperature
		WHERE time_year BETWEEN ? AND ?
		ORDER BY time_year;
	`

	// 执行查询
	rows, err := database.Query(queryStr, start, end)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()
	startYear := start.Year()
	endYear := end.Year()
	years := endYear - startYear + 1
	temp := make([]float64, years)
	i := 0
	for rows.Next() {
		var yeartem float64
		if err := rows.Scan(&yeartem); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			log.Println(err)
			return
		}
		temp[i] = yeartem
		i++
	}
	/*log.Println(endYear)
	log.Println(startYear)
	log.Println("test")
	log.Println(temp)*/
	c.JSON(http.StatusOK, gin.H{
		"yData": temp,
	})
}
