package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type elements struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var vero = elements{"Maria Veronica", 25}
var jose = elements{"Jose Miranda", 30}

var dataName = []elements{vero, jose}

func Pong(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func Name(ctx *gin.Context) {
	nameParam := ctx.Param("name")
	ctx.JSON(http.StatusOK, gin.H{
		"My name": nameParam,
	})
}

func Names(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"users": dataName,
	})
}

func AddUser(ctx *gin.Context) {
	name := ctx.PostForm("name")
	age := ctx.PostForm("age")
	ageInt, _ := strconv.Atoi(age)
	user := elements{name, ageInt}
	ctx.Bind(&user)
	dataName = append(dataName, user)
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Added user",
	})
}
