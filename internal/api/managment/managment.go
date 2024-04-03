package managment

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/niazlv/subscribe-middleware-go/internal/database"
)

func Setup(app *gin.RouterGroup) {
	api := app.Group("api")

	api.POST("subscribe", createSubscribe)
}

func createSubscribe(c *gin.Context) {
	var req database.Subscibe

	db := database.InitDB()

	defer db.Close()

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.Subscribe1 != "" {
		log.Println("req.Subscribe1: " + string(req.Subscribe1))
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Subscribe1 can't be NULL"})
		return
	}
	if req.Subscribe2 != "" {
		log.Println("req.Subscribe2: " + string(req.Subscribe2))
	}
	if req.Id != "" {
		log.Println("req.Id: " + string(req.Id))
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id can't be NULL"})
		return
	}
	if req.Next != "" {
		log.Println("req.Next: " + string(req.Next))
	}

	database.StoreSubscribe(db, req)
	c.JSON(http.StatusOK, gin.H{"id": req.Id})
	//c.Data(http.StatusOK, "text/plain", []byte(""+string(req.Subscribe1)))
}
