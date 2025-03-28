package controllers

import (
	"api_recu_corte1/src/persona/application"
	"api_recu_corte1/src/persona/infrastructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetNewPersonIsAddedController struct {
	app *application.GetNewPersonIsAddedUc
}

func NewGetNewPersonIsAddedController() *GetNewPersonIsAddedController {
	mysql := infrastructure.GetMySQL()
	useCase := application.NewGetNewPersonIsAddedUc(mysql)
	return &GetNewPersonIsAddedController{app: useCase}
}

func (controller *GetNewPersonIsAddedController) Run(ctx *gin.Context) {
	newPersonAdded, err := controller.app.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": false, "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": true, "new_person_added": newPersonAdded})
}
