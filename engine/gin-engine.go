package engine

import "github.com/gin-gonic/gin"

var engine *gin.Engine

func init() {
	engine = gin.Default()
}

//GetAdmin is
func GetAdmin() *gin.RouterGroup {
	return GetEngine().Group("/admin", gin.BasicAuth(gin.Accounts{
		"yoedihar": "2006",
	}))
}

//GetEngine is
func GetEngine() *gin.Engine {
	if engine == nil {
		engine = gin.Default()
	}
	return engine
}
