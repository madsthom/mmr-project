package server

import (
	controllers "mmr/backend/controllers"
	"mmr/backend/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := router.Group("/api/v1")
	{
		mg := v1.Group("/mmr")
		{
			match := new(controllers.MatchController)
			mg.POST("/match", match.SubmitMatch)
		}
		s := v1.Group("/stats")
		{
			leaderboard := new(controllers.LeaderboardController)
			s.GET("/leaderboard", leaderboard.GetLeaderboard)
		}
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router

}
