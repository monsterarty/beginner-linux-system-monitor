package main

import (
	"fmt"
	"os"
	"strings"
	"time"
	"strconv"
	"errors"
	"log/slog"
)


func getHostname() (string, error) {
	file, err := os.Hostname()
	if err != nil {
		slog.Error("os.Hostname error", err)
		return "", err 
	}
	return file, nil
}

func getUptime() (time.Duration, error) {
	file, err := os.ReadFile("/proc/uptime")
	if err != nil {
		slog.Error("os.ReadFile at /proc/uptime", err)
		return 0, err
	}
	uptime, _, verify := strings.Cut(string(file), ".")
	if !verify {
		slog.Error("proc/uptime is empty")
		return 0, errors.New("Empty uptime from file")
	}
	uptimeDur, err := time.ParseDuration(uptime+"s")
	if err != nil {
		slog.Error("Error at uptimeDur - ParsingDuration")
		return 0, errors.New("Error at uptimeDuration")

	}
	return uptimeDur, nil
}

func getCpuInfo() (string, error) {
	file, err := os.ReadFile("/proc/cpuinfo")
	if err != nil {
		slog.Error("Error at ReadFile cpuInfo: ", err)
		return "", err
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
		slog.Error("Error while Cutting cpuInfo for cores")
		return "", errors.New("Empty cpuInfoSuffix")
	}
	return cpuInfoSuf, nil
}

func getKernelVer() (string, error) {
	file, err := os.ReadFile("/proc/version")
	if err != nil {
		slog.Error("Error at ReadFile Version: ", err)
		return "", errors.New("Error at os.ReadFile KernelVer")
	}
	kernel := strings.Split(string(file), " ")
	kernelVer := kernel[0] + " " +kernel[2]
	return kernelVer, nil
}

func helperMemoryInfo(memory string) (float64, error) {
	memoryFloat, err := strconv.Atoi(memory)
	if err != nil {
		slog.Error("Error at helperMemoryInfo > strconv Atoi: ", err)
		return 0.0, errors.New("Error at helperMemoryInfo")
	}
	memoryBack := float64(memoryFloat) / 1000000
	return memoryBack, nil
}

func getMemoryInfo() (float64, float64, error) {
	file, err := os.ReadFile("/proc/meminfo")
	if err != nil {
		slog.Error("Error at ReadFile Meminfo: ", err)
		return 0.0, 0.0, errors.New("Error at os.ReadFile memInfo")	
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
	totalMemoryFloat, err := helperMemoryInfo(totalMemory)
	if err != nil {
		return 0.0, 0.0, err
	}
	freeMemoryFloat, err := helperMemoryInfo(freeMemory)
	if err != nil {
		return 0.0, 0.0, err
	}

	return totalMemoryFloat, freeMemoryFloat, nil
}

func main() {
	fmt.Println("=========================\n Linux System Monitor \n=========================")
	hostname, err := getHostname()
	if err != nil {
		slog.Error("getHostname")
		hostname = "Err"
	}
	kernelVer, err := getKernelVer()
	if err != nil {
		slog.Error("getKernelVer")
		kernelVer = "Err"
	}
	uptime, err := getUptime()
	if err != nil {
		slog.Error("getUptime")
		uptime = 000
	}
	cpuInfo, err := getCpuInfo()
	if err != nil {
		slog.Error("getCpuInfo")
		cpuInfo = "Err"	
	}
	memTotal, memFree, err := getMemoryInfo()
	if err != nil {
		slog.Error("getMemoryInfo")
		memTotal, memFree = 0.0, 0.0
	}
	fmt.Println("Hostname : ", hostname) 
	fmt.Println("Kernel   : ", kernelVer)
	fmt.Println("Uptime   : ", uptime)
	fmt.Println("CPU Cores: ", cpuInfo)
	fmt.Printf("Memory   : %.1f / %.1f Gb", memFree, memTotal)
}
