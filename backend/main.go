package main

import (
	"FinalProject/handler"
	"FinalProject/repository"
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "beasiswa.db")
	if err != nil {
		panic(err)
	}

	Repository := repository.NewRepository(db)
	Handler := handler.NewHandler(*Repository)

	router := gin.Default()
	router.POST("/api/login", Handler.Login)

	router.Run(":8080")
}
