package main

import (
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	_ "net/http"
	"os"
	_ "time"
)

func main() {
	httpRouter := fasthttprouter.New()
	eventController, err := NewEventController()

	if err != nil {
		fmt.Printf("can't create database connection: %s", err)
		os.Exit(1)
	}

	httpRouter.GET("/ping", func(ctx *fasthttp.RequestCtx) {
		response(ctx, "pong", 200)
	})

	httpRouter.GET("/event", eventController.Filter)
	httpRouter.PUT("/event", eventController.Create)
	httpRouter.POST("/event", eventController.Update)
	httpRouter.DELETE("/event", eventController.Delete)

	err = fasthttp.ListenAndServe(":8080", httpRouter.Handler)
	if err != nil {
		fmt.Errorf("can't listen and serve: %w", err)
	}
}