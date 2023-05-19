package main

import (
	"AustemServer/config"
	H "AustemServer/handlers"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.WarnLevel)
}

func main() {
	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}
	//app := server.NewApp()
	//if err := app.Run(viper.GetString("port")); err != nil {
	//	log.Fatalf("%s", err.Error())
	//}
	r := gin.New()
	r.Use(H.CORSMiddleware())

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	err := r.Run(":8080")
	if err != nil {
		return
	}

	fmt.Println("Starting Server at port :8080")
}
