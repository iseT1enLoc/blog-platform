package main

import (
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
	_ = infras.NewAppContext(db, cfg.SecretKey)
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "seminar starting.....")
	})
	r.Run("localhost:8080")
}
