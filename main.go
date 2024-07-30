package main

import (
	"sync"

	"github.com/patrickn2/Golang-Computer/modules/bus"
	"github.com/patrickn2/Golang-Computer/modules/cpu"
	"github.com/patrickn2/Golang-Computer/modules/gpu"
)

// CPU and Video Frequency in Hz
var cpuFrequency int = 10
var gpuFrequency int = 1000

func main() {

	bus := bus.NewBus()
	cpu := cpu.NewCPU(cpuFrequency, bus)
	gpu := gpu.NewVideo(gpuFrequency, bus)
	var wg sync.WaitGroup
	wg.Add(2)

	go cpu.Run()
	go gpu.Run()
	wg.Wait()
}
