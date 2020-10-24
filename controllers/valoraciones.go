package controllers

import (
	"fmt"
	"net/http"
	"reflect"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/PedroMFC/microservicio/models"
	"github.com/gocelery/gocelery"
	"github.com/gomodule/redigo/redis"
)

type CreateValoracionInput struct {
	Asignatura  string `json:"asignatura" binding:"required"`
	Valoracion  int `json:"valoracion" binding:"required"`
}

type MediaAsignaturaOutput struct {
	Asignatura  string `json:"asignatura" binding:"required"`
	Media  int `json:"media" binding:"required"`
}

// GET /valoraciones
// Find all valoraciones
func FindValoraciones(c *gin.Context) {
	redisPool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			  c, err := redis.DialURL("redis://")
			  if err != nil {
				  return nil, err
			  }
			  return c, err
		  },
	  }
	  
	  // initialize celery client
	  cli, _ := gocelery.NewCeleryClient(
		  gocelery.NewRedisBroker(redisPool),
		  &gocelery.RedisCeleryBackend{Pool: redisPool},
		  1,
	  )
	  
	  // prepare arguments
	  taskName := "worker.add"
	  argA := 5
	  argB := 7
	  
	  // run task
	  asyncResult, err := cli.Delay(taskName, argA, argB)
	  if err != nil {
		  panic(err)
	  }
	  
	  // get results from backend with timeout
	  res, err := asyncResult.Get(10 * time.Second)
	  if err != nil {
		  panic(err)
	  }
	  
	  log.Printf("result: %+v of type %+v", res, reflect.TypeOf(res))

	var valoraciones []models.Valoracion
	models.DB.Find(&valoraciones)

	fmt.Println(gin.H{"data": valoraciones})


	c.JSON(http.StatusOK, gin.H{"data": valoraciones})
}

// POST /valoraciones
// Create new valoracion
func CreateValoracion(c *gin.Context) {
	// Validate input
	var input CreateValoracionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	valoracion := models.Valoracion{Asignatura: input.Asignatura, Valoracion: input.Valoracion}
	models.DB.Create(&valoracion)

	c.JSON(http.StatusOK, gin.H{"data": valoracion})
}

// GET /books/:id
// Find a book
func FindValAsignatura(c *gin.Context) {
	// Get model if exist
	var valoraciones[] models.Valoracion
	if err := models.DB.Where("asignatura = ?", c.Param("asignatura")).Find(&valoraciones).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}


	c.JSON(http.StatusOK, gin.H{"data": valoraciones})
}

// GET /books/:id
// Find a book
func FindMediaAsignatura(c *gin.Context) {
		
	// Get model if exist
	var valoraciones[] models.Valoracion
	if err := models.DB.Select("valoracion").Where("asignatura = ?", c.Param("asignatura")).Find(&valoraciones).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var i int = 0
	for _,val:= range valoraciones{
		i = i + val.Valoracion
		log.WithFields(log.Fields{
			"animal": "walrus",
			"size":   10,
		  }).Info("A group of walrus emerges from the ocean: ", val.Valoracion)
	}

	resultado := MediaAsignaturaOutput{Asignatura:  c.Param("asignatura"), Media: i}
	c.JSON(http.StatusOK, gin.H{"data": resultado})
}