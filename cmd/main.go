package main

import (
	"blog-platform-go/api/middlewares"
	"blog-platform-go/infras/appconfig"
	infras "blog-platform-go/infras/appctx"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := appconfig.LoadConfig()

	if err != nil {
		log.Fatalf("There is error while loading config... %s", err)
	}
	fmt.Print("Connecting to database...")
	db, err := appconfig.ConnectDatabaseWithRetryIn20s(cfg)
	if err != nil {
		log.Fatalln("Error when connecting to database:", err)
	}
	appctx := infras.NewAppContext(db, cfg.SecretKey)
	r := gin.Default()

	r.Use(middlewares.CORS())
	r.Use(middlewares.Recover(appctx))

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "passing middlewares successfully")
	})
	r.Run("localhost:8081")
}
