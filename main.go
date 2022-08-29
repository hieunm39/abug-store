package main

import (
	"abug-store/components/appctx"
	"abug-store/helpers"
	"abug-store/middleware"
	"abug-store/modules/auth/authcontroller"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


func main() {
	dsn := helpers.GetDsn("MYSQL_CONNECTION")
	secretKey := helpers.GetSecretKey("SECRET_KEY")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db = db.Debug()

	if err := serve(db, secretKey); err != nil {
		log.Fatalln(err)
	}
}

func serve(db *gorm.DB, secretKey string) error {

	appCtx := appctx.NewAppContext(db, secretKey)

	r := gin.Default()
	r.Use(middleware.Recover(appCtx))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "ok",
		})
	})

	v1 := r.Group("/v1")
	
	auth := v1.Group("/auth")
	{	
		auth.POST("/login", authcontroller.Login(appCtx))
		auth.POST("/register", authcontroller.Register(appCtx))	
	}

	return r.Run(`:8080`)
}