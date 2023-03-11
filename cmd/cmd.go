package cmd

import (
	"fmt"

	"dancer/conf"
)

func Start() {
	fmt.Printf("1 =====> 🚀🚀🚀 %v\n", 1)
	conf.InitConfig()
}

func Clean() {
	fmt.Printf("2 =====> 🚀🚀🚀 %v\n", 2)

}
