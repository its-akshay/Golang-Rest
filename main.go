package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"res.mod/db"
	"res.mod/models"
	"strconv"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent) // all req to following events dynamic
	server.POST("/events", createEvent)
	server.Run(":8080") //local host in which we are listeling request

}
func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError,gin.H{"message":"could not fetch events. Try again later"})
		
		return
	}
	context.JSON(http.StatusOK, events)
}
func getEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"),10,64)
	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"message":"could not parse event id"})
		return 
	}
	event,err:=models.GetEventById(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError,gin.H{"message":"could not fetch event"})
	}
	context.JSON(http.StatusOK, event)

}




func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data."})
		return
	}
	event.ID = 1
	event.UserId = 1


	err  = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError,gin.H{"message":"could not create events. Try again later"})
		
		return
	}
	context.JSON(http.StatusCreated, gin.H{
		"message": "Event created",
		"event":   event,
	})
}