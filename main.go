package main

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/syyfs/test-savechaincode/backend"
)

var configPath = "./config/config.yaml"

func init()  {

	viper.SetConfigFile(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Errorf("===== ReadInConfig faild , err is [%s]======\n", err)
	}
}

func main()  {

	// savechaincode 持续发送交易10个
	token , err := backend.Login()
	if err != nil {
		fmt.Errorf("** Login Faild !!!**\n")
	}

	for i :=1 ; i <=100 ; i++ {
		backend.InvokeSaveValue(token,i)
	}
}
