package main

import (
	"fmt"
	"os"
	"time"

	"github.com/mackerelio/go-osstat/cpu"
	"github.com/mackerelio/go-osstat/memory"
	"github.com/ricochet2200/go-disk-usage/du"
)

var KB = uint64(1024)

func main() {

	before, err := cpu.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}

	time.Sleep(time.Duration(1) * time.Second)
	after, err := cpu.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}
	total := float64(after.Total - before.Total)
	cpu := CPU{
		User:   float64(after.User-before.User) / total * 100,
		System: float64(after.System-before.System) / total * 100,
		Idle:   float64(after.Idle-before.Idle) / total * 100,
	}
	
	memory, err := memory.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}
	mem := Memory{
		Total:  memory.Total,
		Used:   memory.Used,
		Cached: memory.Cached,
		Free:   memory.Free,
	}

	usage := du.NewDiskUsage(".")
	dus := DiskUsage {
		Free: usage.Free(),
		Available: usage.Available(),
		Size: usage.Size(),
		Used: usage.Used(),
		Usage: usage.Usage()*100,
	}

	fmt.Printf("cpu user: %v \n", cpu)
	fmt.Printf("memory total: %v bytes\n", mem)
	fmt.Printf("usage total: %v bytes\n", dus)

	// fmt.Println("Free in MB: ", usage.Free()/(KB*KB))
	// fmt.Println("Available:", usage.Available()/(KB*KB))
	// fmt.Println("Size:", usage.Size()/(KB*KB))
	// fmt.Println("Used:", usage.Used()/(KB*KB))
	// fmt.Println("Usage:", usage.Usage()*100, "%")
}
