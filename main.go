package main

import (
	"crudproject/internal/db"
	"crudproject/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	r := gin.Default()
	routes.SetupRoutes(r, db.DB)
	r.Run(":8080")
}
