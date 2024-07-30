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
		time.Sleep(time.Millisecond * 950 / time.Duration(gpu.frequency))
	}
}

func (gpu *GPU) Clock() {
	if data := gpu.bus.GetDataForGPU(); data != "" {
		go fmt.Println("DATA PRINTED BY GPU", data, time.Now())
	}
}
