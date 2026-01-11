package router

import (
	c "golang/internal/controller"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.GET("/users/:id", c.NewUserController().GetUserByID)
		v1.GET("/companies/:id", c.NewCompanyController().GetCompanyByID)
	}
	return router
}
