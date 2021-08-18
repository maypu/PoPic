package main

import (
	"PoPic/router"
	"PoPic/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// static
	//r = router.Static(r)
	// router
	r = router.Routers(r)
	// run
	runPort := utils.GetConfig("common.port")
	if runPort == "" {
		runPort = "8085"
	}
	err := r.Run(":" + runPort) //run in terminal: fresh
	fmt.Println(err)
}
