package main

import (
	"blog-platform-go/api/middlewares"
	"blog-platform-go/api/routes"
	"blog-platform-go/component/appconfig"
	component "blog-platform-go/component/appctx"
	database "blog-platform-go/infras/postgres"
	"time"

	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := appconfig.LoadConfig()

	if err != nil {
		log.Fatalf("There is error while loading config... %s", err)
	}
	fmt.Print("Connecting to database...")
	db, err := database.ConnectDatabaseWithRetryIn20s(cfg)
	if err != nil {
		log.Fatalln("Error when connecting to database:", err)
	}
	appctx := component.NewAppContext(db, cfg.SecretKey)
	r := gin.Default()

	r.Use(middlewares.CORS())
	r.Use(middlewares.Recover(appctx))

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "passing middlewares successfully")
	})
	routes.Setup(time.Duration(time.Second*20), db, r)
	r.Run("localhost:8081")
}
