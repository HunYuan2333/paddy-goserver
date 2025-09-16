package ShowImg

import (
	"paddy-goserver/DieaseImage"
	"paddy-goserver/GrowImage"
	"paddy-goserver/PredictImage"
	_ "paddy-goserver/PredictImage"
	"paddy-goserver/UserOperation"

	"github.com/gin-gonic/gin"
)

func ShowImg(c *gin.Context) {
	ImgType := c.Param("ImgType")
	switch {
	case ImgType == "PredictDiseaseImage":
		PredictImage.ShowPredictDiseaseImage(c)
	case ImgType == "PredictGrowImage":
		PredictImage.ShowPredictGrowImage(c)
	case ImgType == "GrowImage":
		GrowImage.ShowGrowImage(c)
	case ImgType == "DiseaseImage":
		DieaseImage.ShowDiseaseImage(c)
	case ImgType == "UserImage":
		UserOperation.UserImage(c)

	}
}
