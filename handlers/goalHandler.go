package handlers

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type Goal struct {

	Goal_ID     int       `json:"goal_id"`     
	Title       string    `json:"title"`       
	Description string    `json:"description"` 
	CreatedAt   time.Time `json:"created_at"`  
	UpdatedAt   time.Time `json:"updated_at"`  
	DueDate     time.Time `json:"due_date"`    
	User_ID     int       `json:"user_id"`     

}

// Função para criar uma meta
func (g *Goal) CriarMeta(db *sql.DB) gin.HandlerFunc {

	return func(c *gin.Context) {

		token := c.Request.Header.Get("Authorization")

		userID, err := ValidarOToken(token)

		if err != nil {

			c.JSON(401, gin.H{"message": "Token inválido"})

			return

		}

		var goal Goal

		if err := c.BindJSON(&goal); err != nil {

			fmt.Println(err)

			c.JSON(400, gin.H{"message": "Erro ao criar meta"})

			return

		}

		row := db.QueryRow("INSERT INTO goals (title, description, due_date, user_id) VALUES ($1, $2, $3, $4) RETURNING goal_id, created_at, updated_at", goal.Title, goal.Description, goal.DueDate, userID)

		err = row.Scan(&goal.Goal_ID, &goal.CreatedAt, &goal.UpdatedAt)

		if err != nil {

			fmt.Println(err)

			c.JSON(500, gin.H{"message": "Erro ao criar meta"})

			return

		}

		c.JSON(200, goal)

	}

}

// Função para atualizar uma meta passando o id da meta
func (g *Goal) AtualizarMeta(db *sql.DB) gin.HandlerFunc {

	return func(c *gin.Context) {

		token := c.Request.Header.Get("Authorization")

		userID, err := ValidarOToken(token)

		if err != nil {

			c.JSON(401, gin.H{"message": "Token inválido"})

			return

		}

		goalID := c.Param("id")

		var goal Goal

		if err := c.BindJSON(&goal); err != nil {

			fmt.Println(err)

			c.JSON(400, gin.H{"message": "Erro ao atualizar meta"})

			return

		}

		row := db.QueryRow("UPDATE goals SET title = $1, description = $2 WHERE goal_id = $3 AND user_id = $4 RETURNING updated_at", goal.Title, goal.Description, goalID, userID)

		err = row.Scan(&goal.UpdatedAt)

		if err != nil {

			c.JSON(500, gin.H{"message": "Erro ao atualizar meta"})

			return

		}

		c.JSON(200, goal)

	}

}

// Função para deletar uma meta passando o id da meta
func (g *Goal) DeletarMeta(db *sql.DB) gin.HandlerFunc {
	
	return func(c *gin.Context) {

		token := c.Request.Header.Get("Authorization")

		userID, err := ValidarOToken(token)

		if err != nil {

			c.JSON(401, gin.H{"message": "Token inválido"})

			return

		}

		goalID := c.Param("id")

		_, err = db.Exec("DELETE FROM goals WHERE goal_id = $1 AND user_id = $2", goalID, userID)

		if err != nil {

			c.JSON(500, gin.H{"message": "Erro ao deletar meta"})

			return

		}

		c.JSON(200, gin.H{"message": "Meta deletada com sucesso"})

	}

}

// Função para listar todas as metas do usuario passando o id do usuario
func (g *Goal) ListarMetas(db *sql.DB) gin.HandlerFunc {
	
	return func(c *gin.Context) {

		token := c.Request.Header.Get("Authorization")

		userID, err := ValidarOToken(token)

		if err != nil {

			c.JSON(401, gin.H{"message": "Token inválido"})

			return

		}

		var goals []Goal

		rows, err := db.Query("SELECT goal_id, title, description, created_at, updated_at, due_date, user_id FROM goals WHERE user_id = $1", userID)

		if err != nil {

			c.JSON(500, gin.H{"message": "Erro ao listar metas"})

			fmt.Println(err)

			return

		}

		for rows.Next() {

			var goal Goal

			err := rows.Scan(&goal.Goal_ID, &goal.Title, &goal.Description, &goal.CreatedAt, &goal.UpdatedAt, &goal.DueDate, &goal.User_ID)

			if err != nil {

				c.JSON(500, gin.H{"message": "Erro ao listar metas"})

				return

			}

			goals = append(goals, goal)

		}

		c.JSON(200, gin.H{"goals": goals})

	}

}

// Função para visualizar uma meta passando o id da meta -- Página HTML
func (g *Goal) VisualizarMeta(db *sql.DB) gin.HandlerFunc {

	return func(c *gin.Context) {

		token, _ := c.Cookie("token")

		userID, err := ValidarOToken(token)
		
		if err != nil {
		
			c.JSON(401, gin.H{"message": "Token inválido"})
		
			return
		}

		goalID := c.Param("id")

		var goal Goal

		row := db.QueryRow("SELECT goal_id, title, description, created_at, updated_at, due_date, user_id FROM goals WHERE goal_id = $1 AND user_id = $2", goalID, userID)

		err = row.Scan(&goal.Goal_ID, &goal.Title, &goal.Description, &goal.CreatedAt, &goal.UpdatedAt, &goal.DueDate, &goal.User_ID)
		
		if err != nil {
		
			if err == sql.ErrNoRows {
		
				c.JSON(404, gin.H{"message": "Meta não encontrada"})
		
				return
		
			}
		
			c.JSON(500, gin.H{"message": "Erro ao visualizar meta"})
		
			return
		}

		c.HTML(200, "goal.html", gin.H{"goal": goal})

	}

}
