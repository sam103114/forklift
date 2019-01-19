package main

import (
	"./loadbalancer"
	"time"
)

func main() {
	loadbalancer.ApplySettings()
	ticker := time.NewTicker(time.Duration(loadbalancer.Bal.PeriodicHostCheckDelay) * time.Second)
	go loadbalancer.UpdateAliveHosts(ticker.C)
	loadbalancer.Start()

}
