package controller

import (
	rest_err "github.com/HunCoding/meu-primeiro-crud-go/src/configuration/REST_ERR"
	"github.com/gin-gonic/gin"
)


func CreateUser(c *gin.Context){
	err := rest_err.NewBedRequestError("voce chamou caminho errado")
	c.JSON(err.CODE, err)
}

func FindUserByEmail(c *gin.Context){
}