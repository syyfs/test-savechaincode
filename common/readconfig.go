package common

import "github.com/spf13/viper"

func Getclientcommont()  string {
	return viper.GetString("server.restful.clientcommont")
}


func GetClientPort() string {
	return viper.GetString("server.restful.clientport")
}

func GetChannelName() string {
	return viper.GetString("server.execchaincode.channelname")
}

func GetChaincodeName() string {
	return viper.GetString("server.execchaincode.chaincodename")
}

func GetChaincodeFunc() string {
	return viper.GetString("server.execchaincode.chaincodefunc")
}