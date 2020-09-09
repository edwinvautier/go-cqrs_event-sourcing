package commandbus

import (
	"github.com/gin-gonic/gin"
	"log"
)

func Send(c *gin.Context) {
	log.Println(c.Request.URL.Path)
}