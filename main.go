package main

import "github.com/gin-gonic/gin"

func main() {

	router := gin.Default()

	// Simple group: v1
	// v1 := router.Group("/v1")
	// {
	// 	v1.POST("/login", loginEndpoint)
	// 	v1.POST("/submit", submitEndpoint)
	// 	v1.POST("/read", readEndpoint)
	// }
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	// Simple group: v2
	// v2 := router.Group("/v2")
	// {
	// 	v2.POST("/login", loginEndpoint)
	// 	v2.POST("/submit", submitEndpoint)
	// 	v2.POST("/read", readEndpoint)
	// }

	router.Run(":8080")
}

// Before continue:
//
// Gracefully Shutdown on control+C/command+C or when kill command sent is ENABLED BY-DEFAULT.
//
// In order to manually manage what to do when app is interrupted,
// We have to disable the default behavior with the option `WithoutInterruptHandler`
// and register a new interrupt handler (globally, across all possible hosts).
// func main() {
// 	app := iris.New()

// 	idleConnsClosed := make(chan struct{})
// 	iris.RegisterOnInterrupt(func() {
// 		timeout := 10 * time.Second
// 		ctx, cancel := stdContext.WithTimeout(stdContext.Background(), timeout)
// 		defer cancel()
// 		// close all hosts
// 		app.Shutdown(ctx)
// 		close(idleConnsClosed)
// 	})

// 	app.Get("/", func(ctx iris.Context) {
// 		ctx.HTML(" <h1>hi, I just exist in order to see if the server is closed</h1>")
// 	})

// 	// http://localhost:8080
// 	app.Listen(":8080", iris.WithoutInterruptHandler, iris.WithoutServerError(iris.ErrServerClosed))
// 	<-idleConnsClosed
// }
