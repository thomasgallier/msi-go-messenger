package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type Room struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Messages []Message `json:"messages"`
}

type Message struct {
	ID int `json:"id"`
	Text string `json:"text"`
	Date string `json:"date"`
}

var rooms = map[int]*Room {
	1: {ID: 1, Name: "Charlie", Messages: []Message{}},
	2: {ID: 2, Name: "GuiGui", Messages: []Message{}},
}

func error(c *gin.Context) {
	c.HTML(http.StatusOK, "error.tmpl", gin.H{})
}

func getRoomById(id int) *Room {
	var room Room
	var find, exist = rooms[id]
	if !exist { return &room } else { return find }
}

func homeView(c *gin.Context, room *Room) {
	c.HTML(http.StatusOK, "room.tmpl", gin.H{
		"Room": room,
		"Messages": room.Messages,
		"OnAddMessage" : addMessage,
	})
}

func indexView(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"Rooms": rooms,
	})
}

func home(c *gin.Context) {
	var id, _ = strconv.Atoi(c.Param("room"))
	room := getRoomById(id)
	if room.ID > 0 { homeView(c, room) } else { error(c) }
}


func newMessage(id int, text string) Message {
	return Message{
		ID: id,
		Text: text,
		Date: time.Now().Local().Format("02-01-2006 15:04:05"),
	}
}
func addMessage(c *gin.Context) {
	var id, _ = strconv.Atoi(c.Param("room"))
	room := getRoomById(id)

	if room.ID > 0 {
		text := c.PostForm("text")
		if len(text) > 0 {
			msg := newMessage(len(room.Messages), text)
			room.Messages = append(room.Messages, msg)
		}
		homeView(c, room)
	} else { error(c) }
}

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("views/html/*.tmpl")
	router.Static("/static", "./views/static")

	router.GET("/room/:room", home)
	router.POST("/room/:room", addMessage)

	router.GET("/", indexView)

	router.GET(":12345", error)

	router.Run("localhost:8080")
}

