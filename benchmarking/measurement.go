package benchmarking

import (
	"time"
)

type measurement struct {
	*timer
	opCode     OpCode
	allocation uint32
}

func startNewMeasurement(opCode OpCode) *measurement {
	return &measurement{
		timer:  &timer{startTime: time.Now()},
		opCode: opCode,
	}
}

func (m *measurement) pause() {
	m.stop()
}

func (m *measurement) resume() {
	m.start()
}

func (m *measurement) end(size uint32) {
	m.stop()
	if size != 0 && m.allocation != 0 {
		panic("measurement cannot have both allocation and size")
	} else if size == 0 {
		size = m.allocation
	}

	fileWriter.export(m.opCode, m.elapsedTime, size)
}
