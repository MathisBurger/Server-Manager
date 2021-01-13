package system

import "github.com/shirou/gopsutil/disk"

type Partition struct {
	Device     string  `json:"device"`
	FsType     string  `json:"fs_type"`
	Mountpoint string  `json:"mountpoint"`
	Opts       string  `json:"opts"`
	Used       float64 `json:"used"`
	Free       float64 `json:"free"`
	Total      float64 `json:"total"`
	Percent    float64 `json:"percent"`
}

type StorageInfo struct {
	Num        int         `json:"num"`
	Partitions []Partition `json:"partitions"`
}

func GetStorageInfo() *StorageInfo {
	partitions, _ := disk.Partitions(true)
	main := new(StorageInfo)
	main.Num = len(partitions)
	var parts []Partition
	for _, val := range partitions {
		total, _ := disk.Usage(val.Mountpoint)
		parts = append(parts, Partition{
			val.Device,
			val.Fstype,
			val.Mountpoint,
			val.Opts,
			float64(total.Used) / 1024 / 1024 / 1024,
			float64(total.Free) / 1024 / 1024 / 1024,
			float64(total.Total) / 1024 / 1024 / 1024,
			total.UsedPercent})
	}
	main.Partitions = parts
	return main
}
