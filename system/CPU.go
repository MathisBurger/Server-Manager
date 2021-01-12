package system

import (
	"github.com/shirou/gopsutil/cpu"
	"time"
)

type CPU_Stats struct {
	Workload      []float64 `json:"workload"`
	LogicalCounts int       `json:"logical_counts"`
}

type CPU_Info struct {
	NumCPU     int     `json:"num_cpu"`
	Cores      int     `json:"cores"`
	Name       string  `json:"name"`
	MaxClock   float64 `json:"max_clock"`
	CacheSize  int     `json:"cache_size"`
	CPU_Family string  `json:"cpu_family"`
}

func GetCPU_Stats() CPU_Stats {
	percent, _ := cpu.Percent(time.Duration(1000000000), true)
	logicalCounts, _ := cpu.Counts(true)
	return CPU_Stats{percent, logicalCounts}
}

func GetCPU_Info() *CPU_Info {
	info, _ := cpu.Info()
	strct := new(CPU_Info)
	strct.NumCPU = len(info)
	strct.Cores = int(info[0].Cores)
	strct.Name = info[0].ModelName
	strct.MaxClock = info[0].Mhz
	strct.CacheSize = int(info[0].CacheSize)
	strct.CPU_Family = info[0].Family
	return strct
}
