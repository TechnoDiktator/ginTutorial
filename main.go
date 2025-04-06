package main

import (
	//"io"
	// "net/http"
	// "time"

	//"github.com/TechnoDiktator/ginTutorial/middleware"

	"io"
	"os"

	"github.com/sirupsen/logrus"

	//"github.com/TechnoDiktator/ginTutorial/middleware"
	//"github.com/TechnoDiktator/ginTutorial/middleware/logger"
	"github.com/gin-gonic/gin"
	//"google.golang.org/grpc/admin"
)

func getData(c *gin.Context) {

	c.JSON(200, gin.H{
		"data": "Hi I am getData GIN Framework",
	})

}

func main() {

	// router := gin.Default()
	logrus.SetLevel(logrus.TraceLevel)
	logrus.Traceln("Trace level log")
	logrus.Debugln("Debug level log")
	logrus.Infoln("Info level log")
	logrus.Warnln("Warn level log")
	logrus.Errorln("Error level log")
	logrus.Fatalln("Fatal level log")
	logrus.Panicln("Panic level log")
	logrus.SetReportCaller(true)

	f, _ := os.Create("logrus.log")
	multi := io.MultiWriter(f, os.Stdout)

	logrus.SetOutput(multi)

	router := gin.New()
	logrus.Println("Starting the server...")
	router.GET("/getData", getData)

	router.Run(":8080")

}
