package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("Error at Hostname: ", err)
	}
	
	file, err := os.ReadFile("/proc/uptime")
	if err != nil {
		fmt.Println("Error at ReadFile: ", err)
	}
	uptimeRaw, _, verify := strings.Cut(string(file), " ")
	if !verify {
		fmt.Println("Error at uptime string Cut, nothing to parse")
	}
	uptimeRawSec, _, verify := strings.Cut(uptimeRaw, ".")
	if !verify {
		fmt.Println("Error at float point string Cut, nothing to Cut")
	}
	uptime, err := time.ParseDuration(uptimeRawSec+"s")
	if err != nil {
		fmt.Println("Error at ParsingDuration uptimeSec", err)
	}
	
	fmt.Println("=========================\n Linux System Monitor \n=========================")
	fmt.Println("Hostname: ", hostname) 
	fmt.Println("Kernel   : TODO")
	fmt.Println("Uptime   : ", uptime)
	fmt.Println("\nCPU Cores: TODO")
	fmt.Println("Memory     : proc/meminfo")
}
