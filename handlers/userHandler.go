package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type User struct {

	User_ID   int       `json:"user_id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	FullName  string    `json:"full_name"`
	CreatedAt time.Time `json:"created_at"`

}

// Função para logar o usuário
func (u *User) Entrar(db *sql.DB) gin.HandlerFunc {

	return func(c *gin.Context) {

		var user User

		if err := c.BindJSON(&user); err != nil {

			c.JSON(400, gin.H{"message": "Erro ao fazer login"})

			return

		}

		row := db.QueryRow("SELECT user_id, username, password FROM users WHERE username = $1 AND password = $2", user.Username, user.Password)

		err := row.Scan(&user.User_ID, &user.Username, &user.Password)

		if err != nil {

			c.JSON(404, gin.H{"message": "Usuário ou senha incorretos"})

			return

		}

		// Gere o token genérico
        token, err := GerarOToken(user)

        if err != nil {
            c.JSON(500, gin.H{"message": "Erro ao gerar token"})
            return
        }

		cookie := &http.Cookie{

			Name: "token",

			Value: token,

			Expires:  time.Now().Add(72 * time.Hour),

			HttpOnly: false,

			Secure: false,

			SameSite: http.SameSiteStrictMode,

			Path: "/",
		}

		http.SetCookie(c.Writer, cookie)

		c.JSON(200, gin.H{"message": "Login efetuado com sucesso!", "token": token})

	}

}

// Função para registrar o usuário
func (u *User) Registrar(db *sql.DB) gin.HandlerFunc {

	return func(c *gin.Context) {

		var newUser User

		if err := c.BindJSON(&newUser); err != nil {

			c.JSON(400, gin.H{"message": "Erro ao criar usuario"})

			return

		}

		_, err := db.Exec("INSERT INTO users (username, email, password, full_name, created_at) VALUES ($1, $2, $3, $4, $5)", newUser.Username, newUser.Email, newUser.Password, newUser.FullName, time.Now())

		fmt.Println(newUser.Username, newUser.Email, newUser.Password, newUser.FullName, c.ClientIP(), time.Now())

		if err != nil {

			c.JSON(500, gin.H{"message": "Erro ao criar usuário"})

			return

		}

		c.JSON(200, gin.H{"message": "Usuário criado com sucesso!"})

	}

}

// Função para deletar o usuário
func (u *User) Deletar(db *sql.DB) gin.HandlerFunc {
	
	return func(c *gin.Context) {

		// Lendo o atual token

		token := c.Request.Header.Get("Authorization")

		// Verificando se o token é válido

		userID, err := ValidarOToken(token)

		if err != nil {

			c.JSON(401, gin.H{"message": "Token inválido"})

			return

		}

		_, err = db.Exec("DELETE FROM users WHERE user_id = $1", userID)

		if err != nil {

			c.JSON(404, gin.H{"message": "Usuário não encontrado"})

			return

		}

		cookie := &http.Cookie{

			Name: "token",

			Value: "",

			Expires: time.Unix(0, 0),

			HttpOnly: true,

			Secure: true,

			SameSite: http.SameSiteStrictMode,

		}

		http.SetCookie(c.Writer, cookie)

		c.JSON(200, gin.H{"message": "Usuário deletado com sucesso!"})

	}

}