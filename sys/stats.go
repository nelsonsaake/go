package sys

import (
	"sort"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
	"github.com/shirou/gopsutil/v3/process"
)

type ProcessInfo struct {
	Name string  `json:"name"`
	CPU  float64 `json:"cpu"`
	Mem  float32 `json:"mem"`
}

type Tasks struct {
	Running  int `json:"running"`
	Sleeping int `json:"sleeping"`
	Stopped  int `json:"stopped"`
	Zombie   int `json:"zombie"`
}

type SystemStats struct {
	CPUPercent []float64 `json:"cpu_percent"`

	CPUTimesUser   float64 `json:"cpu_times_user"`
	CPUTimesSystem float64 `json:"cpu_times_system"`

	MemoryUsedPercent float64 `json:"memory_used_percent"`
	SwapUsedPercent   float64 `json:"swap_used_percent"`

	Load1  float64 `json:"load_1"`
	Load5  float64 `json:"load_5"`
	Load15 float64 `json:"load_15"`

	Uptime uint64 `json:"uptime"`

	Tasks Tasks `json:"tasks"`

	DiskUsedPercent float64 `json:"disk_used_percent"`

	NetBytesSent uint64 `json:"net_bytes_sent"`
	NetBytesRecv uint64 `json:"net_bytes_recv"`

	ProcessCount int `json:"process_count"`

	TopProcesses []ProcessInfo `json:"top_processes"`
}

func GetSystemStats(args ...int) (*SystemStats, error) {

	topN := 15
	if len(args) > 0 {
		topN = args[0]
	}

	s := &SystemStats{}

	// CPU
	cpuPerc, _ := cpu.Percent(0, false)
	cpuTimes, _ := cpu.Times(false)

	// Memory / Swap
	vm, _ := mem.VirtualMemory()
	sm, _ := mem.SwapMemory()

	// Load / Host
	l, _ := load.Avg()
	h, _ := host.Info()

	// Disk / Network
	d, _ := disk.Usage("/")
	n, _ := net.IOCounters(false)

	// Processes
	procs, _ := process.Processes()

	s.CPUPercent = cpuPerc
	s.CPUTimesUser = cpuTimes[0].User
	s.CPUTimesSystem = cpuTimes[0].System

	s.MemoryUsedPercent = vm.UsedPercent
	s.SwapUsedPercent = sm.UsedPercent

	s.Load1 = l.Load1
	s.Load5 = l.Load5
	s.Load15 = l.Load15

	s.Uptime = h.Uptime

	s.DiskUsedPercent = d.UsedPercent

	s.NetBytesSent = n[0].BytesSent
	s.NetBytesRecv = n[0].BytesRecv

	s.ProcessCount = len(procs)

	var list []ProcessInfo

	for _, p := range procs {
		name, _ := p.Name()
		cpuP, _ := p.CPUPercent()
		memP, _ := p.MemoryPercent()

		list = append(list, ProcessInfo{
			Name: name,
			CPU:  cpuP,
			Mem:  memP,
		})
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].CPU > list[j].CPU
	})

	if topN > len(list) {
		topN = len(list)
	}
	s.TopProcesses = list[:topN]

	// task states
	var r, sl, st, z int
	for _, p := range procs {
		states, _ := p.Status()
		if len(states) == 0 {
			continue
		}
		switch states[0] {
		case "R":
			r++
		case "S":
			sl++
		case "T":
			st++
		case "Z":
			z++
		}
	}

	s.Tasks = Tasks{
		Running:  r,
		Sleeping: sl,
		Stopped:  st,
		Zombie:   z,
	}

	return s, nil
}
