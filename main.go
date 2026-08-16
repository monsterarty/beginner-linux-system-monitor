package main

import (
	"fmt"
	"os"
	"strings"
	"time"
	"strconv"
)

func getHostname() string {
	file, err := os.Hostname()
	if err != nil {
		fmt.Println("Error at Hostname generation: ", err)
	}
	return file 
}

func getUptime() time.Duration {
	file, err := os.ReadFile("/proc/uptime")
	if err != nil {
		fmt.Println("Error at osReadFile proc/uptime :", err)
	}
	uptime, _, verify := strings.Cut(string(file), ".")
	if !verify {
		fmt.Println("proc/uptime is empty")
	}
	uptimeDur, err := time.ParseDuration(uptime+"s")
	if err != nil {
		fmt.Println("Error at uptimeDur - ParsingDuration")
	}
	return uptimeDur 
}

func getCpuInfo() string {
	file, err := os.ReadFile("/proc/cpuinfo")
	if err != nil {
		fmt.Println("Error at ReadFile cpuInfo: ", err)
	}
	cpuInfoLines := strings.Lines(string(file))
	var cpuInfo string
	for c := range cpuInfoLines {
		prefix, found := strings.CutPrefix(c, "siblings")
		if found {
			cpuInfo = prefix
			break
		}
	}
	_, cpuInfoSuf, verify := strings.Cut(cpuInfo, " ") 
	if !verify {
		fmt.Println("Error while Cutting cpuInfo for cores")
	}
	return cpuInfoSuf
}

func getKernelVer() string {
	file, err := os.ReadFile("/proc/version")
	if err != nil {
		fmt.Println("Error at ReadFile Version: ", err)
	}
	kernel := strings.Split(string(file), " ")
	kernelVer := kernel[0] + " " +kernel[2]
	return kernelVer
}

func helperMemoryInfo(memory string) float64 {
	memoryFloat, err := strconv.Atoi(memory)
	if err != nil {
		fmt.Println("Error at helperMemoryInfo > strconv Atoi: ", err)
	}
	memoryBack := float64(memoryFloat) / 1000000
	return memoryBack 
}

func getMemoryInfo() (float64, float64) {
	file, err := os.ReadFile("/proc/meminfo")
	if err != nil {
		fmt.Println("Error at ReadFile Meminfo: ", err)
	}
	memoryLines := strings.Split(string(file), "\n")
	var totalMemory string
	var freeMemory string

	for _, memoryLine := range memoryLines {
		if len(memoryLine) == 0 {
			continue
		}
		lineParts := strings.Fields(memoryLine)
		if lineParts[0] == "MemTotal:" {
			totalMemory = lineParts[1]
		}
		if lineParts[0] == "MemFree:" {
			freeMemory = lineParts[1]
		}
		if len(totalMemory) > 1 && len(freeMemory) > 1 {
			break
		}
	}	
	totalMemoryFloat := helperMemoryInfo(totalMemory)
	freeMemoryFloat := helperMemoryInfo(freeMemory)

	return totalMemoryFloat, freeMemoryFloat
}

func main() {
	fmt.Println("=========================\n Linux System Monitor \n=========================")
	fmt.Println("Hostname : ", getHostname()) 
	fmt.Println("Kernel   : ", getKernelVer())
	fmt.Println("Uptime   : ", getUptime())
	fmt.Println("CPU Cores: ", getCpuInfo())
	memTotal, memFree := getMemoryInfo()
	fmt.Printf("Memory   : %.1f / %.1f Gb", memFree, memTotal)
}
