package configuration

import (
    "encoding/json"
    "os"
    "fmt"
)

type Settings struct {
   GlobalDebug string
   DBdebug string
   DBhost string
   DBname string
   DBuser string
   DBpass string
   SSLCERT string
   SSLKEY string
   HTTPPORT string
   SSLPORT string
   Programworkingdirectory string

}

var Gdebug bool

func init() {
   Gdebug = false
}

func Loadsettings(conf *Settings) {
   file, _ := os.Open("/etc/viral/conf.json")
   decoder := json.NewDecoder(file)
   err := decoder.Decode(&conf)
   if err != nil {
      fmt.Println("configuration module: Error in configuration file:", err)
   } else {
      if (conf.GlobalDebug == "1") {
         Gdebug = true
         fmt.Println("configuration module: config file loaded!")
      }
   }
}

