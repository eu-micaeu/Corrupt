package routes

import (
	
	"database/sql"

	"github.com/gin-gonic/gin"

	"github.com/eu-micaeu/Corrupt/handlers"

)

func GoalRoutes(r *gin.Engine, db *sql.DB) {
	
	goalHandler := handlers.Goal{}

	r.POST("/createGoal", goalHandler.CriarMeta(db))

	r.GET("/listGoals", goalHandler.ListarMetas(db))

	r.DELETE("/deleteGoal/:id", goalHandler.DeletarMeta(db))

	r.GET("/viewGoal/:id", goalHandler.VisualizarMeta(db))

}