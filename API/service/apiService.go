package service

import (
	"log"
	"net/http"

	"API/model"

	"github.com/gin-gonic/gin"
)

type NewService struct {
	Name    string  `json:"name" binding:"required"`
	Purpose string  `json:"purpose" binding:"required"`
	SLA     float64 `json:"sla" binding:"required"`
	Price   float64 `json:"price" binding:"required"`
}

type ServiceUpdate struct {
	Name    string  `json:"name"`
	Purpose string  `json:"purpose"`
	SLA     float64 `json:"sla"`
	Price   float64 `json:"price"`
}

func GetServices(c *gin.Context) {
	var services []model.Service

	db, err := model.Database()
	if err != nil {
		log.Println(err)
	}

	if err != nil {
		db.Find(&services)
		return
	}
	c.JSON(http.StatusOK, services)
}

func GetService(c *gin.Context) {
	var service model.Service

	db, err := model.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Where("id= ?", c.Param("id")).First(&service).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Service not found"})
		return
	}

	c.JSON(http.StatusOK, service)
}

func PostService(c *gin.Context) {
	var service NewService

	if err := c.ShouldBindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newService := model.Service{Name: service.Name, Price: service.Price, Purpose: service.Purpose, SLA: service.SLA}

	db, err := model.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Create(&newService).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, newService)
}

func UpdateService(c *gin.Context) {

	var service model.Service

	db, err := model.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Where("id = ?", c.Param("id")).First(&service).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Service not found!"})
		return
	}

	var updateService ServiceUpdate

	if err := c.ShouldBindJSON(&updateService); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Model(&service).Updates(model.Service{Name: service.Name, Price: service.Price, Purpose: service.Purpose, SLA: service.SLA}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, service)

}

func DeleteService(c *gin.Context) {

	var service model.Service

	db, err := model.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Where("id = ?", c.Param("id")).First(&service).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Service not found!"})
		return
	}

	if err := db.Delete(&service).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Service deleted"})

}