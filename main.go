package main

import (
	"api_recu_corte1/src/core"
	person "api_recu_corte1/src/persona/infrastructure"
	routesPerson"api_recu_corte1/src/persona/infrastructure/routes"

	"github.com/gin-gonic/gin"
)



func main (){
	person.GoMySQL()

	r:= gin.Default()

	core.InitCORS(r)

	routesPerson.Routes(r)

	r.Run()
}