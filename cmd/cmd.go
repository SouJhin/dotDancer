package cmd

import (
	"fmt"

	"dancer/conf"
	"dancer/router"
)

func Start() {
	fmt.Printf("1 =====> 🚀🚀🚀 %v\n", 1)
	conf.InitConfig()
	router.InitRouter()
}

func Clean() {
	fmt.Printf("2 =====> 🚀🚀🚀 %v\n", 2)
}
