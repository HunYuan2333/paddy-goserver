package ShowImg

import (
	"github.com/gin-gonic/gin"
	"paddy-goserver/DieaseImage"
	"paddy-goserver/GrowImage"
	"paddy-goserver/PredictImage"
	_ "paddy-goserver/PredictImage"
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
		//case ImgType==

	}
}
