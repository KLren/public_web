package routers

import (
	"website/server/controllers"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	//r.Use(controllers.CORSMiddleWare())

	grouprV1 := r.Group("api/v1")
	{
		grouprV1.GET("/products", controllers.GetALLProductsFromDB)
		grouprV1.GET("/products/:id", controllers.GetProductByID)
		grouprV1.GET("/users", controllers.GetALLUsersFromDB)
		grouprV1.GET("/users/:id", controllers.GetUserByID)
		grouprV1.POST("/users/login", controllers.LoginUser)
		//grouprV1.POST("/products/:id", ModifyProductByID)
	}
}
