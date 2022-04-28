package Routes

import (
	"first/Controllers"
	"first/Middleware"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes() *gin.Engine {

	r := gin.Default()

	r.POST("/login", Controllers.Login)
	r.GET("/health", Controllers.HealthCheck)
	r.POST("/register", Controllers.SaveUser)

	authorized := r.Group("/", Middleware.VerifyToken)
	{
		articlePrefix := authorized.Group("/articles")
		{
			articlePrefix.GET("/all", Controllers.GetArticles)
			articlePrefix.GET("/article/:id", Controllers.GetArticleById)
			articlePrefix.GET("/comments/:id", Controllers.GetCommentsForArticle)
		}
		commentPrefix := authorized.Group("/articles")
		{
			commentPrefix.POST("/comment", Controllers.SaveComment)
		}

	}
	return r
}
