package sys

import (
	"math"
	"time"
	"github.com/shirou/gopsutil/cpu"
	"github.com/wailsapp/wails"
)

// User Stats, warning and whatnot 

// Stats .
type Stats struct {
	log *wails.CustomLogger
}

// CPUUsage .
type CPUUsage struct {
	Average int `json:"avg"`
}

// Num Cores 
type CPUDiskNum struct {
	Disks int `json:"disk"`
}
// WailsInit .
func (s *Stats) WailsInit(runtime *wails.Runtime) error {
	s.log = runtime.Log.New("Stats")

	go func() {
		for {
			runtime.Events.Emit("cpu_usage", s.GetCPUUsage())
			time.Sleep(1 * time.Second)
		}
	}()

	return nil
}


// Times
func (s *Stats) GetTimes() []cpu.TimesStat {
	diskInfo, err := cpu.Times(true)
	if err != nil {
		s.log.Errorf("unable to get cpu times: %s", err.Error())
		return nil
	}
	return diskInfo
}

// GetCPUUsage .
func (s *Stats) GetCPUUsage() *CPUUsage {
	percent, err := cpu.Percent(1*time.Second, false)
	if err != nil {
		s.log.Errorf("unable to get cpu stats: %s", err.Error())
		return nil
	}

	return &CPUUsage{
		Average: int(math.Round(percent[0])),
	}
}

func (s *Stats) GetDiskSerialNum(logical bool) *CPUDiskNum {
	diskNum, err := cpu.Counts(logical)
	if err != nil {
		s.log.Errorf("unable to get cpu stats: %s", err.Error())
		return nil
	}
	return &CPUDiskNum{
		Disks: int(diskNum),
	}
}

func (s *Stats) GetInfo() []cpu.InfoStat {
	diskInfo, err := cpu.Info()
	if err != nil {
		s.log.Errorf("unable to get cpu info: %s", err.Error())
		return nil
	}
	return diskInfo
}