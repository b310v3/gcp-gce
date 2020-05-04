package main

import (
	"fmt"

	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"

	"github.com/shirou/gopsutil/cpu"
)

func main() {
	v, _ := mem.VirtualMemory()
	c, _ := cpu.Info()
	h, _ := host.Info()

	// almost every return value is a struct
	fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)

	// convert to JSON. String() is also implemented
	fmt.Println(v)

	fmt.Println(c)

	fmt.Println(h)
}
