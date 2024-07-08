package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"testmod/config"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := config.LoadEnv(".env"); err != nil {
		unwrapped := errors.Unwrap(err)
		fmt.Printf("%T\n", unwrapped)
	}

	r := gin.Default()

	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.JSON(http.StatusOK, message)
	})
	r.GET("/checkers", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})
	routerErr := r.Run(":8080")
	if routerErr != nil {
		log.Fatal(routerErr)
	}
}
