package gpu

import (
	"fmt"
	"time"

	"github.com/patrickn2/Golang-Computer/modules/bus"
)

type GPU struct {
	frequency int
	bus       *bus.BUS
}

func NewVideo(frequency int, bus *bus.BUS) *GPU {
	return &GPU{
		frequency: frequency,
		bus:       bus,
	}
}

func (gpu *GPU) Run() {
	for {
		gpu.Clock()
		time.Sleep(time.Second / time.Duration(gpu.frequency))
	}
}

func (gpu *GPU) Clock() {
	fmt.Println("GPU Tick", gpu.frequency, time.Now())
	if data := gpu.bus.GetDataForGPU(); data != "" {
		fmt.Println("DATA PRINTED BY GPU", data)
	}
}
