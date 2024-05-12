package main

import (
	"database/sql"

	"github.com/gin-contrib/cors"
	storage "github.com/niazlv/subscribe-middleware-go/internal/database"
	"github.com/niazlv/subscribe-middleware-go/internal/routes"

	"github.com/gin-gonic/gin"
)

var db *sql.DB

func main() {
	router := gin.Default()

	router.Use(cors.Default())

	routes.Setup(router)

	db = storage.InitDB()
	storage.CreateTable(db)

	router.Run(":56792")
}
