package bookhizhan

import (
	"fmt"
	"testing"
)

type ChipType int

const (
	None ChipType =iota
	CPU // 中央处理器
	GPU // 图形处理器
)

func (c ChipType) String() string{
	switch c{
	case None:
		return "None"
	case CPU:
		return "CPU"
	case GPU:
		return "GPU"
	}
	return "N/A"
}
func TestIota(t *testing.T){
	fmt.Printf("%s %d",CPU,CPU)
}

