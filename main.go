package main

import (
	"github.com/chengziwj/mudan/controller"
	"github.com/chengziwj/mudan/infrastructure/persistence"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	services, err := persistence.NewRepositories("root", "v1nDZ3OhzRKZLCFr", 3306, "127.0.0.2", "mudan")
	if err != nil {
		log.Fatal(err)
	}

	users := controller.NewUserController(services.User)

	r := gin.Default()
	r.POST("/api/user", users.SaveUser)
	r.GET("/api/user/:uid", users.GetUser)
	r.Run(":8080")
}
