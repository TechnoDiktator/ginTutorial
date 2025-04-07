package main

import ( //"io"

	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/http2"

	//"google.golang.org/grpc/admin"

	"github.com/TechnoDiktator/ginTutorial/serviceinit"
)

func main() {
	// Initialize all required services
	localServiceInit, err := serviceinit.LocalServerInit()

	if err != nil {
		logrus.Error("Error while initializing required services: %w", err)
		return
	}
	// Ensure that the PostgreSQL clients are closed when the function exits to prevent resource leaks.
	defer localServiceInit.PgConnect.Close()

	// // ================================  HTTP Block Starts  =====================================
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "healthy"})
	})

	router.POST("/test", func(c *gin.Context) {
		var requestBody map[string]interface{}
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			return
		}
		logrus.Infof("Received request: %v", requestBody)
		c.JSON(http.StatusOK, gin.H{"status": "success"})
	})

	router.POST("/temp", func(c *gin.Context) {
		controller.RouteRequestHandleTemp(c, localServiceInit)
	})

}

func startServer() {
	server := &http.Server{
		Addr:              ":8080",
		WriteTimeout:      15 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
	}

	err := http2.ConfigureServer(server, nil)

	if err != nil {
		logrus.Errorf("Error configuring HTTP/2: %v", err)
		return
	}

	go func() {
		err := server.ListenAndServe()

		if err != nil {
			logrus.Errorf("Error starting server: %v", err)
		}
	}()

	GracefullShutdown(server)

}

func GracefullShutdown(server *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logrus.Errorf("Server forced to shutdown: %v", err)
	}

	logrus.Infof("Server exiting")
}
