package controllers

import (
	"api_recu_corte1/src/persona/application"
	"api_recu_corte1/src/persona/infrastructure"
	"net/http"

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
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": "Valor de sexo inválido. Debe ser 'true' o 'false'."})
		return
	}

	count, err := controller.app.Execute(sexoBool)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "error": err.Error()})
		return
	}

	if count == 0 {
		c.JSON(http.StatusNoContent, gin.H{"status": true, "message": "No hay registros con el valor de sexo proporcionado."})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": true, "conteo": count})
	}
}
