package system

import "github.com/shirou/gopsutil/mem"

type MemoryInfo struct {
	RAM  *RAM_Info  `json:"ram"`
	SWAP *Swap_Info `json:"swap"`
}

type RAM_Info struct {
	TotalMemory uint64 `json:"total_memory"`
	Buffers     uint64 `json:"buffers"`
	CommitLimit uint64 `json:"commit_limit"`
}

type Swap_Info struct {
	Total uint64 `json:"total"`
}

type MemoryStats struct {
	RAM  *RAM_Stats  `json:"ram"`
	SWAP *Swap_Stats `json:"swap"`
}

type RAM_Stats struct {
	Used    uint64  `json:"used"`
	Free    uint64  `json:"free"`
	Percent float64 `json:"percent"`
}

type Swap_Stats struct {
	Used    uint64  `json:"used"`
	Free    uint64  `json:"free"`
	Percent float64 `json:"percent"`
}

func GetMemory_Stats() *MemoryStats {
	ram, _ := mem.VirtualMemory()
	RAM := new(RAM_Stats)
	main := new(MemoryStats)

	RAM.Used = ram.Used / 1024 / 1024
	RAM.Free = ram.Free / 1024 / 1024
	RAM.Percent = ram.UsedPercent
	main.RAM = RAM

	swap, _ := mem.SwapMemory()
	SWAP := new(Swap_Stats)
	SWAP.Used = swap.Used / 1024 / 1024
	SWAP.Free = swap.Free / 1024 / 1024
	SWAP.Percent = swap.UsedPercent
	main.SWAP = SWAP
	return main
}

func GetMemoryInfo() *MemoryInfo {
	ram, _ := mem.VirtualMemory()
	main := new(MemoryInfo)
	RAM := new(RAM_Info)
	RAM.TotalMemory = ram.Total / 1024 / 1024
	RAM.Buffers = ram.Buffers
	RAM.CommitLimit = ram.CommitLimit
	main.RAM = RAM

	swap, _ := mem.SwapMemory()
	SWAP := new(Swap_Info)
	SWAP.Total = swap.Total / 1024 / 1024
	main.SWAP = SWAP
	return main
}
