package main

import (
	"database/sql"

	storage "github.com/niazlv/subscrbe-middleware-go/internal/database"
	"github.com/niazlv/subscrbe-middleware-go/internal/routes"

	"github.com/gin-gonic/gin"
)

var secretPath = "/noncesub/"

var db *sql.DB

func main() {
	router := gin.Default()

	routes.Setup(router)

	db = storage.InitDB()
	storage.CreateTable(db)

	router.Run(":56792")
}
