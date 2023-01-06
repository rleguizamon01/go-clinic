package main

import (
	"common"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
	"log"
	"medicines"
	"net/http"
)

var Mel *melody.Melody

func main() {
	common.ConnectToDatabase()
	router := newRouter()
	err := router.Run("127.0.0.1:8080")

	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func newRouter() *gin.Engine {
	r := gin.Default()
	Mel = common.CreateWebSocket()

	r.GET("/medicines", medicines.GetMedicinesHandler)
	r.GET("/medicines/:id", medicines.GetMedicineHandler)
	r.POST("/medicines", medicines.CreateMedicineHandler)
	r.PUT("/medicines/:id", medicines.UpdateMedicineHandler)
	r.DELETE("/medicines/:id", medicines.DeleteMedicineHandler)

	r.GET("/medicine-logs", func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, "assets/medicine-logs.html")
	})

	r.GET("/ws", func(c *gin.Context) {
		err := Mel.HandleRequest(c.Writer, c.Request)
		if err != nil {
			log.Panic(err.Error())
		}
	})

	Mel.HandleMessage(func(s *melody.Session, msg []byte) {
		err := Mel.Broadcast(msg)
		if err != nil {
			log.Panic(err.Error())
		}
	})
	return r
}
