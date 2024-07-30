package cpu

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/patrickn2/Golang-Computer/modules/bus"
)

type CPU struct {
	frequency int
	ticker    *time.Ticker
	bus       *bus.BUS
	reader    *bufio.Scanner
	clocks    int
}

func NewCPU(frequency int, bus *bus.BUS) *CPU {
	return &CPU{
		frequency: frequency,
		ticker:    time.NewTicker(time.Second / time.Duration(frequency)),
		bus:       bus,
		reader:    bufio.NewScanner(os.Stdin),
		clocks:    0,
	}
}

func (cpu *CPU) Run() {
	for {
		select {
		case <-cpu.ticker.C:
			cpu.Clock()
		}
	}
}

func (cpu *CPU) Clock() {
	fmt.Println("-CPU Tick", cpu.frequency, cpu.clocks)
	if cpu.clocks != 0 {
		cpu.clocks--
		return
	}

	cpu.ExecInstruction()
}

func (cpu *CPU) ExecInstruction() {
	cpu.clocks = 50
	cpu.bus.SendDataToBus("GPU", "Hello, I'm the CPU talking to you. I'm Gonna execute again after 50 CPU Ticks")
}
