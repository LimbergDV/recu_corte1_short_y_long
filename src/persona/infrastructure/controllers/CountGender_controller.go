package controllers

import (
	"api_recu_corte1/src/persona/application"
	"api_recu_corte1/src/persona/infrastructure"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type CountGenderController struct {
	app *application.CountGenderUc
}

func NewCountGenderController() *CountGenderController {
	mysql := infrastructure.GetMySQL()
	useCase := application.NewCountGenderUc(mysql)
	return &CountGenderController{app: useCase}
}

func (controller *CountGenderController) Run(c *gin.Context) {
	sexo := c.Param("sexo")

	var sexoBool bool
	if sexo == "true" {
		sexoBool = true
	} else if sexo == "false" {
		sexoBool = false
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Valor de sexo inválido. Debe ser 'true' o 'false'."})
		return
	}
	timeout := time.After(30 * time.Second)

	checkInterval := time.NewTicker(5 * time.Second)
	defer checkInterval.Stop()

	for {
		select {
		case <-timeout:
			c.JSON(http.StatusRequestTimeout, gin.H{"error": "Long polling timed out"})
			return
		case <-checkInterval.C:
			count, err := controller.app.Execute(sexoBool)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			if count > 0 { 
				c.JSON(http.StatusOK, gin.H{"conteo": count})
				return
			}
		}
	}

}
