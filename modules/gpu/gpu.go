package gpu

import (
	"fmt"
	"time"

	"github.com/patrickn2/Golang-Computer/modules/bus"
)

type GPU struct {
	frequency int
	ticker    *time.Ticker
	bus       *bus.BUS
}

func NewVideo(frequency int, bus *bus.BUS) *GPU {
	return &GPU{
		frequency: frequency,
		bus:       bus,
		ticker:    time.NewTicker(time.Second / time.Duration(frequency)),
	}
}

func (gpu *GPU) Run() {
	for {
		select {
		case <-gpu.ticker.C:
			gpu.Clock()
		}
	}
}

func (gpu *GPU) Clock() {
	fmt.Println("GPU Tick", gpu.frequency)
	if data := gpu.bus.GetDataForGPU(); data != "" {
		fmt.Println("DATA PRINTED BY GPU", data)
	}
}
