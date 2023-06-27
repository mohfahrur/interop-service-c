package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	googleD "github.com/mohfahrur/interop-service-c/domain/google"
	entity "github.com/mohfahrur/interop-service-c/entity"
	ticketUC "github.com/mohfahrur/interop-service-c/usecase/ticket"
)

func main() {

	// load .env file from given path
  	// we keep it empty it will load .env from current directory
  	err := godotenv.Load(".env")

  	if err != nil {
    	log.Fatalf("Error loading .env file")
  	}


	log.SetFlags(log.Llongfile)
	spreadsheetID := os.Getenv("spreadsheetID")

	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
		return
	}

	credentialsFile, err := ioutil.ReadFile(pwd + "/credential.json")
	if err != nil {
		log.Println(err)
		return
	}

	googleDomain := googleD.NewGoogleDomain(credentialsFile, spreadsheetID)
	ticketUsecase := ticketUC.NewTicketUsecase(*googleDomain)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong from service c",
		})
		return
	})
	r.POST("/update-data", func(c *gin.Context) {
		var req entity.UpdateSheetRequest
		err := c.BindJSON(&req)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "bad request",
			})
			return
		}
		err = ticketUsecase.UpdateSheet(req)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "bad request",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
		return
	})
	r.Run(":5002")
}
