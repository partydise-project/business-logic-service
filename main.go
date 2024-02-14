package main

import (
	"business-logic-service/database"
	"business-logic-service/server"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func initializeEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	initializeEnv()

	dns := os.Getenv("DNS")

	err := database.InicializeDB(dns)
	if err != nil {
		fmt.Println("Error when initializing the database: ", err)
		return
	}

	r := server.InicializeRouter()
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "pending")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT")
		//c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		c.Next()
	})

	puerto := 8080
	direction := fmt.Sprintf(":%d", puerto)

	fmt.Println("Hello server listing in:", puerto)
	log.Fatal(r.Run(direction))
}
