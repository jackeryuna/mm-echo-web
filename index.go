package main

import (
	"fmt"
	"flag"
	"mm-echo-web/router"
)

var helpInfo string = "help\n -h 帮助\n -c conf/conf.toml 配置文件路径，默认conf/conf.toml\n"
var cmdConf = flag.Bool("c", false, "配置文件路径")
var cmdHelp = flag.Bool("h", false, "帮助")

func main() {
	var confFilePath string = ""
	flag.Parse()
	if *cmdConf {
		if flag.NArg() == 1 {
			confFilePath = flag.Arg(0)
			fmt.Printf("run with conf:%s\n", confFilePath)
		} else {
			fmt.Println("-c 参数错误\n" + helpInfo)
			return
		}
	} else if *cmdHelp {
		fmt.Println(helpInfo)
		return
	}

	router.RunSubdomains(confFilePath)
}