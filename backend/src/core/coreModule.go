package core

import (
		"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/ufoscout/Codership/backend/src/configuration"

	"context"
	"log"
	"net"
	"net/http"
	"time"
	"strconv"
)

type coreModule struct {
	config     *configuration.Config
	ginRouter  *gin.Engine
	httpServer *http.Server
	port       int
}

func CoreModule(config *configuration.Config) *coreModule {
	module := coreModule{}

	module.config = config
	module.ginRouter = gin.Default()

	return &module
}

func (c *coreModule) Server() *gin.Engine {
	return c.ginRouter
}

func (c *coreModule) Start() {

	log.Printf("Loading static resources from %s\n", c.config.Server.ResourcesPath)
	c.ginRouter.Use(static.Serve("/", static.LocalFile(c.config.Server.ResourcesPath, true)))

	log.Printf("Starting Server at requested port %d\n", c.config.Server.Port)

	listener, err := net.Listen("tcp", ":" + strconv.Itoa(c.config.Server.Port))
	if err != nil {
		log.Fatal("Error:", err)
	}

	c.port = listener.Addr().(*net.TCPAddr).Port

	c.httpServer = &http.Server{
		Handler: c.ginRouter,
	}

	log.Printf("Starting Server at real port %d\n", c.port)

	c.httpServer.Serve(listener)

}

func (c *coreModule) ServerPort() int {
	return c.port
}

func (c *coreModule) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := c.httpServer.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
}
