package main

import (
	"log"
	"os"

	"example.com/go-seat-finder/router"
	"github.com/gin-gonic/gin"
)

func main() {
	gin := gin.Default()

	router.Router(gin)

	err := gin.Run(os.Getenv("PORT"))
	if err != nil {
		log.Fatal(err)
	}

}
