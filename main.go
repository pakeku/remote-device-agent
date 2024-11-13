package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image/png"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/kbinani/screenshot"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
)

func main() {
	fmt.Println("Remote Device Agent Starting...")

	for {
		checkSystemMetrics()
		captureScreenshot()         // Take screenshot periodically
		time.Sleep(5 * time.Second) // Adjust as needed for regular data collection
	}
}

// Struct to format system metrics for JSON
type SystemMetrics struct {
	CPUUsage    float64 `json:"cpu_usage"`
	MemoryUsage float64 `json:"memory_usage"`
	DiskUsage   float64 `json:"disk_usage"`
	TotalMemory uint64  `json:"total_memory_mb"`
	FreeMemory  uint64  `json:"free_memory_mb"`
	TotalDisk   uint64  `json:"total_disk_gb"`
	FreeDisk    uint64  `json:"free_disk_gb"`
}

func checkSystemMetrics() {
	cpuPercent, _ := cpu.Percent(0, false)
	vmStat, _ := mem.VirtualMemory()
	diskStat, _ := disk.Usage("/")

	metrics := SystemMetrics{
		CPUUsage:    cpuPercent[0],
		MemoryUsage: vmStat.UsedPercent,
		DiskUsage:   diskStat.UsedPercent,
		TotalMemory: vmStat.Total / 1024 / 1024,
		FreeMemory:  vmStat.Free / 1024 / 1024,
		TotalDisk:   diskStat.Total / 1024 / 1024 / 1024,
		FreeDisk:    diskStat.Free / 1024 / 1024 / 1024,
	}

	sendMetrics(metrics)
}

// Capture a screenshot of the primary display and save as PNG
func captureScreenshot() {
	bounds := screenshot.GetDisplayBounds(0)
	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		log.Println("Error capturing screenshot:", err)
		return
	}

	file, err := os.Create("screenshot.png")
	if err != nil {
		log.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	png.Encode(file, img)
	log.Println("Screenshot captured and saved.")
}

// Function to send metrics to a remote server
func sendMetrics(metrics SystemMetrics) {
	jsonData, err := json.Marshal(metrics)
	if err != nil {
		log.Println("Error marshaling JSON:", err)
		return
	}

	url := "http://your-server-url/api/metrics" // Replace with your server endpoint

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("Error sending data:", err)
		return
	}
	defer resp.Body.Close()

	log.Println("Metrics sent successfully, status:", resp.Status)
}
