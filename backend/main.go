package main

import (
	"net/http"

	docs "mmr/backend/docs"
	mmr "mmr/backend/mmr"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@BasePath	/api/v1

// @Summary Submit a match
// @Description Submit a match for MMR calculation
// @Accept json
// @Produce json
// @Param match body mmr.Match true "Match object"
// @Success 200 {string} string "match submitted"
// @Router /mmr/match [post]
func SubmitMatch(c *gin.Context) {
	var json mmr.Match
	err := c.BindJSON(&json)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "match submitted")
}

func main() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		eg := v1.Group("/mmr")
		{
			eg.POST("/match", SubmitMatch)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8080")
}
