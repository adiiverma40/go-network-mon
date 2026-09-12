package main

import (
	"fmt"

	"github.com/adiiverma40/Go_Network_Mon/pkg/network"
)






func main(){
	fmt.Println("-------------G0 Network-Mon!----------------")

	isCaptive, err := network.CheckPortal()
	if err != nil {
			fmt.Printf("Status: %v\n", err)
			return
		}

		if isCaptive {
			fmt.Println("Captive portal detected! Triggering auto-login sequence...")
			// TODO: Execute your login logic here
		} else {
			fmt.Println("Internet is connected. Sleeping until next check.")
		}
	latency, err := network.PingCloudFlare()
	if err != nil {
    fmt.Println("Error:", err)
	} else {
    fmt.Printf("Internet is connected! Ping took %.2f milliseconds.\n", latency)
	}

	protal_err :=  network.ReadPortalInfo()

	if protal_err != nil {
		fmt.Println("err  : %v ", protal_err)
	}
}
