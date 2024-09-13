package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/eu-micaeu/Corrupt/database"
	"github.com/eu-micaeu/Corrupt/middlewares"
	"github.com/eu-micaeu/Corrupt/routes"
)

// main function
func main() {

	r := gin.Default()

	r.Use(middlewares.CorsMiddleware())

	db, err := database.NewDB()

	if err != nil {

		panic(err)

	}

	r.LoadHTMLGlob("./views/*.html")

	r.GET("/", func(c *gin.Context) {

		c.HTML(http.StatusOK, "index.html", nil)

	})

	r.GET("/home", func(c *gin.Context) {

		c.HTML(http.StatusOK, "home.html", nil)

	})

	r.GET("/register", func(c *gin.Context) {

		c.HTML(http.StatusOK, "register.html", nil)

	})

	r.GET("/config", func(c *gin.Context) {

		c.HTML(http.StatusOK, "config.html", nil)

	})

	r.Static("/static", "./static")

	routes.UserRoutes(r, db)

	routes.GoalRoutes(r, db)

	r.Run()
	
}