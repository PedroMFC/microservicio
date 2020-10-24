package main

import (
	"github.com/PedroMFC/microservicio/controllers"
	"github.com/PedroMFC/microservicio/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Connect to database
	models.ConnectDatabase()

	// Routes
	r.GET("/books", controllers.FindBooks)
	r.GET("/books/:id", controllers.FindBook)
	// curl -X POST -H "Content-Type: application/json" -d '{ "title": "Start with Why", "author": "Simon Sinek" }' http://localhost:8080/books
	r.POST("/books", controllers.CreateBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	r.GET("/valoraciones", controllers.FindValoraciones)
	r.GET("/valoraciones/:asignatura", controllers.FindValAsignatura)
	r.GET("media/:asignatura", controllers.FindMediaAsignatura)
	// curl -X POST -H "Content-Type: application/json" -d '{ "asignatura": "TDA", "valoracion": 3 }' http://localhost:8080/valoraciones
	// curl -X POST -H "Content-Type: application/json" -d '{ "asignatura": "TDA", "valoracion": 5 }' http://localhost:8080/valoraciones
	// curl -X POST -H "Content-Type: application/json" -d '{ "asignatura": "CC", "valoracion": 1 }' http://localhost:8080/valoraciones
	r.POST("/valoraciones", controllers.CreateValoracion)

	// Run the server
	r.Run()
}

// CREATE ROLE microservicio WITH LOGIN PASSWORD 'microservicio';
// CREATE DATABASE valoraciones OWNER microservicio;