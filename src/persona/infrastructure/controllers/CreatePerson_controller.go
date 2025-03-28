package controllers

import (
	"api_recu_corte1/src/persona/application"
	"api_recu_corte1/src/persona/domain"
	"api_recu_corte1/src/persona/infrastructure"
	"api_recu_corte1/src/persona/infrastructure/routes/validators"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreatePersonController struct {
	app *application.CreatePerson
}

func NewCreatePersonController() *CreatePersonController {
	mysql := infrastructure.GetMySQL()
	app := application.NewCreatePerson(mysql)
	return &CreatePersonController{app: app}
}

func (cp_c *CreatePersonController) Run (c *gin.Context){
	var persons domain.Person

	if err := c.ShouldBindJSON(&persons); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": "Datos inválidos" + err.Error()})
		return
	}

	if err := validators.CheckPerson(persons); err != nil {
		c.JSON(http.StatusBadRequest, gin.H {"status": false, "error": "Datos invalidos" + err.Error()})
	}
	
	rowsAffected, err := cp_c.app.Run(persons)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	if rowsAffected == 0{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusCreated, gin.H {"mensaje": "Empleado creado"})
		c.JSON(http.StatusOK, persons)
	}
}