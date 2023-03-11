package conf

import (
	"fmt"

	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetConfigName("settings")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./conf/")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf(" init failed =====> 🚀🚀🚀 %v\n", err)
	}
	print(viper.GetString("server.port"))
}
