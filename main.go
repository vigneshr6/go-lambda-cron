package main

import (
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var stage string

type mydata struct {
	Name string `json:"name"`
}

func init() {
	stage = os.Getenv("stage")
}

func main() {
	if stage == "prd" {
		lambda.Start(cronHandler)
	} else {
		startHttpServer()
	}
}

func cronHandler(event events.CloudWatchEvent) error {
	log.Info(event.Resources)
	return nil
}

func startHttpServer() {
	app := gin.Default()
	app.POST("/", httpHandler)

	log.Fatal(app.Run(":8080"))
}

func httpHandler(ctx *gin.Context) {
	var m mydata
	if err := ctx.ShouldBindJSON(&m); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}
	process(m)
	ctx.Status(http.StatusOK)
}

func process(m mydata) {
	log.Info("msg : ", m)
}
