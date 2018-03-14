package kale

import (
	"net/http"
	"log"
	"encoding/json"
	"time"
	"github.com/shirou/gopsutil/mem"
	_"github.com/shirou/gopsutil/disk"
)


type MachineInfo struct{
	ID string
	Timestamp       time.Time
	Total_mem uint64
	Free_mem uint64
	Used_mem uint64
	Total_disk uint64
	Free_disk uint64
	Used_disk uint64

}

func poll(){
	//log.Print("Polled!")
	v, _ := mem.VirtualMemory()
	var mi = MachineInfo{
		ID: ID.String(),
		Total_mem: v.Total,
		Free_mem: v.Free,
		Used_mem: v.UsedPercent
	}
	//var td = HB{ID: "00020023001023"}
	json.NewEncoder(*w).Encode(mi)
}	