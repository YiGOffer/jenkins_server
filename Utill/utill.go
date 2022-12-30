package Utill

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"jenkinsServer/Model"
	"log"

	"gopkg.in/ini.v1"
)
var (
	configPath string
	Conf Model.MainConfig
)

//初始化配置
func InitConfig() {
	cfg, err := ini.Load("Config/env.ini")
	if(err != nil) {
		// 后续优化写入日志文件，方便定位问题
		log.Panicln("load config error!",err.Error())
	}
	env := cfg.Section("env").Key("environment").String()
	if env == "local" {
		configPath = "Config/config_test.json"
		fmt.Println("init local config")
	} else {
		configPath = "Config/config.json"
	}

	buf, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Panicln("load config conf failed: ", err)
	}
	err = json.Unmarshal(buf, &Conf)
	if err != nil {
		log.Panicln("decode config file failed:", string(buf), err)
	}
	fmt.Println("InitConfig Success!\n")
}
