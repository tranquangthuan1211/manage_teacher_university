package main

import (
	"fmt"
	"myapi/api"
	"myapi/database"
	"myapi/docs"
	"myapi/utils"
	"runtime"

	"github.com/gin-gonic/gin"
)

func configRuntime() {
	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)
	fmt.Println("Number of CPUs: ", numCPU)
}
func main() {
	configRuntime()
	var db, err = database.Migration()
	if err != nil {
		fmt.Println("Can not migrate the database! - ", err)
	} else {
		fmt.Println("Migrate the database successfully!")
	}
	ip := utils.GetOutboundIP()
	if utils.RUNNING_MODE == gin.ReleaseMode {
		docs.SwaggerInfo.Host = "api.drl.vn"
	} else {
		docs.SwaggerInfo.Host = ip + ":8080"
	}
	fmt.Println(db)
	fmt.Println("Hello, World!")
	api.RunServer(db)
}
