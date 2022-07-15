package main

import (
	"log"

	"go_tests/model"
	"go_tests/service"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := model.Database()
	if err != nil {
		log.Println(err)
	}
	db.DB()

	router := gin.Default()

	router.GET("/services", service.GetServices)
	router.GET("/services/:id", service.GetService)
	router.POST("/services", service.PostService)
	router.PUT("/service/:id", service.UpdateService)
	router.DELETE("/service/:id", service.DeleteService)

	log.Fatal(router.Run(":10000"))
}
