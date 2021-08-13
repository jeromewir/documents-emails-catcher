package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jeromewir/invoices-fwder/config"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	err := config.ReadFromEnvironment()

	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	dc := NewDropboxClient()

	mailHandler := NewMailHandler(dc)

	r.POST("/api/emails", mailHandler.HandleIncomingEmail)

	r.Run()
}
