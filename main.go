package main

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"main.go/routes"
	"main.go/utils"
)

var (
	DB  *sql.DB
	err error
)

func main() {
	r := gin.Default()

	utils.ConnectDatabase()

	routes.AuthRoutes(r)
	routes.ExpenseRoutes(r)
	routes.CategoryRoutes(r)
	routes.BudgetRoutes(r)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the Personal Finance Management API - BerUang",
		})
	})

	r.Run(":" + os.Getenv("PORT"))

}
