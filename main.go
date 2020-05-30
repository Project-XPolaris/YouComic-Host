package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/static", "./static")
	r.Run(":8885") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}