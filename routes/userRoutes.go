package routes

import (
	
	"database/sql"

	"github.com/gin-gonic/gin"

	"github.com/eu-micaeu/Corrupt/handlers"

)

func UserRoutes(r *gin.Engine, db *sql.DB) {
	
	userHandler := handlers.User{}

	r.POST("/login", userHandler.Entrar(db))

	r.POST("/register", userHandler.Registrar(db))

	r.GET("/rescue", userHandler.Resgatar(db))

	r.DELETE("/delete", userHandler.Deletar(db))

}