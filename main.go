package main

import (
	"fmt"
	"myapi/api"
	"myapi/database"
	"runtime"
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
	fmt.Println(db)
	fmt.Println("Hello, World!")
	api.RunServer(db)
}
