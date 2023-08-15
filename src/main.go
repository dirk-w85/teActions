package main

import (
	"fmt"
	"log"
//	"net/http"
//	"io/ioutil"
//	"encoding/json"
//	"os"
//	"strings"
	//"flag"
//	"time" 
//	"strconv"
	"github.com/spf13/viper"
)

func Logger(msg string){
	if viper.GetBool("global.debug"){
		log.Println(msg)
	}
}


func main()  {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("toml") // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")               // optionally look for config in the working directory
	viper.AddConfigPath("../")               // optionally look for config in the working directory
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	if viper.GetBool("global.debug") {
		Logger("Debugging enabled")
	}

	var teOauthToken = viper.GetString("thousandeyes.oauthToken")
	var teUser = viper.GetString("thousandeyes.user")

	Logger("ThousandEyes Oauth Token: "+teOauthToken)
	Logger("ThousandEyes User: "+teUser)

}