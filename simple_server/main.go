package main

import (
	"sandbox-go/simple_server/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	//router.LoadHTMLGlob("views/*.html")
	//router.Static("/assets", "./assets")

	//store := sessions.NewRedisStore(0)
	//router.Use(sessions.StartDefaultSession(store))

	//user := router.Group("/user")
	//{
	//	user.POST("/signup", routes.UserSignUp)
	//	user.POST("/login", routes.UserLogIn)
	//}

	router.GET("/", routes.Home)
	//router.GET("/login", routes.LogIn)
	//router.GET("/signup", routes.SignUp)
	router.NoRoute(routes.NoRoute)

	router.Run(":8080")
}
