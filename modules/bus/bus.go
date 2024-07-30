package bus

type BUS struct {
	CPUData string
	GPUData string
	APUData string
}

func NewBus() *BUS {
	return &BUS{
		CPUData: "",
		GPUData: "",
		APUData: "",
	}
}

func (b *BUS) SendDataToBus(address, data string) {
	if address == "GPU" {
		b.GPUData = data
		return
	}
	if address == "CPU" {
		b.CPUData = data
		return
	}
	if address == "APU" {
		b.APUData = data
		return
	}
}

func (b *BUS) GetDataForGPU() string {
	gpuData := b.GPUData
	b.GPUData = ""
	return gpuData
}
