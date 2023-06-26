package main

import (
	"database/sql"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mohfahrur/interop-service-c/domain/database"
	googleD "github.com/mohfahrur/interop-service-c/domain/google"
	entity "github.com/mohfahrur/interop-service-c/entity"
	"github.com/mohfahrur/interop-service-c/middleware"
	ticketUC "github.com/mohfahrur/interop-service-c/usecase/ticket"
	userUC "github.com/mohfahrur/interop-service-c/usecase/user"
)

func main() {
	log.SetFlags(log.Llongfile)

	spreadsheetID := os.Getenv("spreadsheetID")
	dbConfig := os.Getenv("dbConfig")

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
	db, err := sql.Open("mysql", dbConfig)
	if err != nil {
		log.Println("Failed to connect to the database:", err)
		return
	}
	defer db.Close()

	userDomain := database.NewUserDomain(db)
	googleDomain := googleD.NewGoogleDomain(credentialsFile, spreadsheetID)
	ticketUsecase := ticketUC.NewTicketUsecase(*googleDomain)
	userUsecase := userUC.NewUserUsecase(*userDomain)

	r := gin.Default()
	r1 := r.Group("/v1")
	r1.Use(middleware.AuthAndAuthorize(*userUsecase))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong from service c",
		})
		return
	})
	r1.GET("/get-users/:id", func(c *gin.Context) {
		datas, err := userUsecase.UserDomain.GetUserInfo(c.Param("id"))
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "bad request",
			})
			return
		}
		c.JSON(http.StatusOK, datas)
		return
	})

	r1.GET("/get-users", func(c *gin.Context) {
		datas, err := userUsecase.UserDomain.GetUsers()
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "bad request",
			})
			return
		}
		c.JSON(http.StatusOK, datas)
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
