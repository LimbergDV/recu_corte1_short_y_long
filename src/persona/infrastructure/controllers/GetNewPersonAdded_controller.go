package controllers

import (
	"api_recu_corte1/src/persona/application"
	"api_recu_corte1/src/persona/infrastructure"
	"net/http"
	"time"

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
	pollingInterval := 5 * time.Second
	timeout := time.After(30 * time.Second)

	ticker := time.NewTicker(pollingInterval)
	defer ticker.Stop()


	for {
		select {
		case <-timeout:
			ctx.JSON(http.StatusRequestTimeout, gin.H{"error": "Polling timed out"})
			return
		case <-ticker.C:
			newPersonAdded, err := controller.app.Execute()
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			if newPersonAdded {
				ctx.JSON(http.StatusOK, gin.H{"new_person_added": true})
				return
			}
		}
	}
}

