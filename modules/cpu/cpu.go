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
	bus       *bus.BUS
	reader    *bufio.Scanner
	clocks    int
}

func NewCPU(frequency int, bus *bus.BUS) *CPU {
	return &CPU{
		frequency: frequency,
		bus:       bus,
		reader:    bufio.NewScanner(os.Stdin),
		clocks:    1,
	}
}

func (cpu *CPU) Run() {
	for {
		cpu.Clock()
		time.Sleep(time.Duration(time.Millisecond * 950 / time.Duration(cpu.frequency)))
	}
}

func (cpu *CPU) Clock() {
	go fmt.Println("CPU Tick", time.Now())
	if cpu.clocks != 1 {
		cpu.clocks--
		return
	}

	cpu.ExecInstruction()
}

func (cpu *CPU) ExecInstruction() {
	cpu.clocks = 20
	cpu.bus.SendDataToBus("GPU", fmt.Sprintf("Hello, I'm the CPU talking to you. I'm Gonna execute again after %d CPU Ticks", cpu.clocks))
}
