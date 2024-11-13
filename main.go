package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
)

func main() {
	fmt.Println("Remote Device Agent Starting...")

	for {
		checkSystemMetrics()
		time.Sleep(5 * time.Second) // Adjust as needed for regular data collection
	}
}

func checkSystemMetrics() {
	// CPU Usage
	cpuPercent, _ := cpu.Percent(0, false)
	fmt.Printf("CPU Usage: %.2f%%\n", cpuPercent[0])

	// Memory Usage
	vmStat, _ := mem.VirtualMemory()
	fmt.Printf("Memory Usage: %.2f%% (Total: %v MB, Free: %v MB)\n",
		vmStat.UsedPercent, vmStat.Total/1024/1024, vmStat.Free/1024/1024)

	// Disk Usage
	diskStat, _ := disk.Usage("/")
	fmt.Printf("Disk Usage: %.2f%% (Total: %v GB, Free: %v GB)\n",
		diskStat.UsedPercent, diskStat.Total/1024/1024/1024, diskStat.Free/1024/1024/1024)
}
