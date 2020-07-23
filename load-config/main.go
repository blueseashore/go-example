package main

import (
	"fmt"
	"github.com/micro/go-micro/v2/config"
)
type Host struct {
	Address string `json:"address"`
	Port    int    `json:"port"`
}

type Config struct {
	Hosts map[string]struct {
		Address string `json:"address"`
		Port    int    `json:"port"`
	} `json:"hosts"`
}

func main() {
	_ = config.LoadFile("/tmp/config.json")

	//var conf Config

	//_ = config.Scan(&conf)

	// 10.0.0.1 3306
	//fmt.Println(conf)
	//fmt.Println(conf.Hosts["cache"].Address, conf.Hosts["database"].Port)
	//fmt.Println(config.Get("hosts", "database", "port").Int(3000))
	//fmt.Println(config.Get("hosts", "database", "port2").Int(3000))

	w, err := config.Watch("hosts", "database")
	if err != nil {
		fmt.Println("failed to watch :", err)
	}
	v, err := w.Next()
	if err != nil {
		fmt.Println("failed to next:", err)
	}

	var host Host
	fmt.Println(v)
	_ = v.Scan(&host)
	fmt.Println(host)
}
