/*
 * @Descripttion: integrated to initialize
 * @Author:DW
 * @Date: 2024-04-24 16:32:41
 * @LastEditTime: 2024-04-24 18:15:12
 */
package api

import (
	"cassandra/src/cmd/utils"
	"cassandra/src/pkg/client/cassandra"
	"cassandra/src/pkg/handler"
	"cassandra/src/pkg/repository/db"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Initialize(config utils.Configuration) {
	router := gin.Default()
	fmt.Printf("confgi is %+v\n", config)

	session := cassandra.ConnectDatabase(config.Database.Url, config.Database.Keyspace)

	rep := db.NewTodoRepository(session)
	orderHandler := handler.NewTodoHandler(&rep)
	router.GET("/ping", orderHandler.HealthCheck)

	router.POST("api/v1/todo/", orderHandler.CreateTodo)
	router.GET("api/v1/todo/:id", orderHandler.GetTodoById)

	//run the server :8080
	router.Run(":" + config.Host.Port + "")
}
