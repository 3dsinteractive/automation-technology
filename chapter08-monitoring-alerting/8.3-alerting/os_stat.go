package main

import (
	"os"
	"time"

	"github.com/mackerelio/go-osstat/cpu"
	"github.com/mackerelio/go-osstat/memory"
)

type OSStat struct {
	Hostname            string `json:"host_name"`
	MemoryTotal         int    `json:"mem_total_mbytes"`
	MemoryUsed          int    `json:"mem_used_mbytes"`
	MemoryCached        int    `json:"mem_cached_mbytes"`
	MemoryFree          int    `json:"mem_free_mbytes"`
	MemoryUsedPercent   int    `json:"mem_used_percent"`
	MemoryCachedPercent int    `json:"mem_cached_percent"`
	MemoryFreePercent   int    `json:"mem_free_percent"`
	CPUTotal            int    `json:"cpu_total"`
	CPUUserPercent      int    `json:"cpu_user_percent"`
	CPUSystemPercent    int    `json:"cpu_system_percent"`
	CPUIdlePercent      int    `json:"cpu_idle_percent"`
}

// osStat will block 1 second to process the stat of CPU (and other metric)
func osStat() (*OSStat, error) {

	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}

	memory, err := memory.Get()
	if err != nil {
		return nil, err
	}

	before, err := cpu.Get()
	if err != nil {
		return nil, err
	}

	time.Sleep(time.Duration(1) * time.Second)
	after, err := cpu.Get()
	if err != nil {
		return nil, err
	}

	memTotal := int(memory.Total / (1024 * 1024))
	memUsed := int(memory.Used / (1024 * 1024))
	memCached := int(memory.Cached / (1024 * 1024))
	memFree := int(memory.Free / (1024 * 1024))
	cpuTotal := int(after.Total - before.Total)

	osStat := &OSStat{
		Hostname:            hostname,
		MemoryTotal:         memTotal,
		MemoryUsed:          memUsed,
		MemoryCached:        memCached,
		MemoryFree:          memFree,
		MemoryUsedPercent:   int(float64(memory.Used) / float64(memory.Total) * 100),
		MemoryCachedPercent: int(float64(memory.Cached) / float64(memory.Total) * 100),
		MemoryFreePercent:   int(float64(memory.Free) / float64(memory.Total) * 100),
		CPUTotal:            int(cpuTotal),
		CPUUserPercent:      int(float64(after.User-before.User) / float64(cpuTotal) * 100),
		CPUSystemPercent:    int(float64(after.System-before.System) / float64(cpuTotal) * 100),
		CPUIdlePercent:      int(float64(after.Idle-before.Idle) / float64(cpuTotal) * 100),
	}
	return osStat, nil
}
