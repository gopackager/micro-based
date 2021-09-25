package config

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/gopackager/micro-based/transports/middlewares"
)

type Router interface {
	RegisterRoutes(opt []Routing) *router
	Starts()
}

type router struct {
	appEnv           string
	appVersionRoutes string
	appPort          string
	server           *http.Server
	engine           *gin.Engine
}

func New(app Application) Router {
	switch app.Mode {
	case "production", "staging":
		gin.SetMode(gin.ReleaseMode)
	case "test":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.DebugMode)
	}
	return &router{
		appEnv:           app.Mode,
		appVersionRoutes: app.Version,
		appPort:          app.Port,
		server:           &http.Server{},
		engine:           gin.New(),
	}
}

func (r *router) RegisterRoutes(optsHandler []Routing) *router {
	r.engine.Use(
		gin.Recovery(),
		gzip.Gzip(gzip.DefaultCompression),
		middlewares.AppID,
		middlewares.RequestID,
	)

	route := r.engine.Group(r.appVersionRoutes)
	{
		for _, v := range optsHandler {
			switch strings.ToLower(v.Method) {
			case "get":
				route.GET(v.Path, v.Handler)
			case "post":
				route.POST(v.Path, v.Handler)
			case "put":
				route.PUT(v.Path, v.Handler)
			case "delete":
				route.DELETE(v.Path, v.Handler)
			}
		}
	}

	r.server = &http.Server{
		Addr:    fmt.Sprintf(":%v", r.appPort),
		Handler: r.engine,
	}
	return r
}

func (r router) Starts() {
	go func() {
		if err := r.server.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down services...")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := r.server.Shutdown(ctx); err != nil {
		log.Fatal("services forced to shutdown:", err)
	}

	log.Println("services exiting")
}
