package routers

import (
	"fmt"
	"golang-gin-boilerplate/serverConfig"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Routers interface {
	RunRouter()
}

type router struct {
	config serverConfig.Config
	db     *gorm.DB
}

func Init(config serverConfig.Config, db *gorm.DB) *router {
	return &router{config, db}
}

func (r *router) RunRouter() {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	apiV1 := router.Group("/api/v1")

	userRouters := CreateUserRouter(r, apiV1)
	userRouters.InitRouter()

	err := router.Run(fmt.Sprintf(":%s", r.config.Port))
	if err != nil {
		panic("[Error] failed to start Gin server due to: " + err.Error())
	}
}
