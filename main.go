package main

import (
	"github.com/batthanhvan/middlewares"
	"github.com/batthanhvan/src/controllers"
	"github.com/batthanhvan/src/db"
	"github.com/batthanhvan/src/lib"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db.ConnectDataBase()
	r := gin.Default()
	r.SetTrustedProxies(nil)

	// gin.SetMode(gin.ReleaseMode)
	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "PUT", "POST", "DELETE"},
		AllowHeaders:  []string{"Authorization", "Content-Type", "User-Agent"},
		ExposeHeaders: []string{"content-disposition", "content-description"},
	}))

	// api := r.Group("/api")
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.GET("/user", controllers.CurrentUser)

	// admin := r.Group("/admin")
	// admin.Use(middlewares.Only(lib.ROLE_ADMIN))

	reportGroup := r.Group("/reports")
	reportGroup.Use(middlewares.Only(lib.ROLE_ADMIN))
	reportGroup.GET("", controllers.HandleShowAllReports)
	reportGroup.GET("/:username", controllers.HandleGetReportByUsername)

	playerGroup := r.Group("/players")
	playerGroup.GET("/:username", controllers.HandleGetByUserName)
	// playerGroup.POST("/restrict/:username", controllers.HandlePostRestrictPlayerTime)
	playerGroup.POST("/modify/:playername/:tagline", controllers.HandlePostModifyUser)

	matchGroup := r.Group("/matches")
	matchGroup.GET("/:username", controllers.HandleGetByMatchID)

	r.Run("localhost:8080")
}
