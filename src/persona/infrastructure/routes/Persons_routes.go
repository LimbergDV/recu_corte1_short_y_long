package routes

import (
	"api_recu_corte1/src/persona/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func Routes (r *gin.Engine) {
	
	employeesRoutes := r.Group("/persons") 
	{
		employeesRoutes.POST("/", controllers.NewCreatePersonController().Run)

	}
}