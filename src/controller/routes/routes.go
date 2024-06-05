package routes

import (
	"github.com/HunCoding/meu-primeiro-crud-go/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes (r * gin.RouterGroup){
	r.GET("./getUserByID/:userID", controller.FindUserById)
	r.GET("./getUserEmail/:userID", controller.FindUserByEmail)
	r.POST("./createUser", controller.CreateUser)
	r.PUT("./updateUser/:userID", controller.Updateuser)
	r.DELETE("./geleteuser/:userID", controller.DeleteUser)
	
        }