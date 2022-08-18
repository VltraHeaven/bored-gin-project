package main

import (
	distro2 "bored-gin-project/api/v1/distro"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1", "localhost"})
	router.GET("/v1/distros", distro2.GetDistros)
	router.POST("/v1/distros", distro2.PostDistro)
	router.GET("/v1/distros/:id", distro2.GetDistroByID)
	router.Run(":8080")
}
