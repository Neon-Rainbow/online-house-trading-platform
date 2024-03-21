package root

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetUpRootAPI(router *gin.Engine, db *gorm.DB) {
	router.GET("/", HomePageGet)
	router.GET("/learn_more", LearnMoreGet)
}
