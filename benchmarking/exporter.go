package benchmarking

import (
	"encoding/binary"
	"fmt"
	"os"
	"time"
)

const flushTimerInterval = time.Duration(time.Second * 30)

var fileWriter *exporter

func initExporter(fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		panic("could not create benchmark file: " + err.Error())
	}

	fileWriter = &exporter{
		file:              file,
		bytesToFlushAfter: 10 * 100000,
		flushTimer:        *time.NewTimer(flushTimerInterval),
	}

	go func() {
		for {
			<-fileWriter.flushTimer.C
			fileWriter.file.Sync()
			fileWriter.flushTimer.Reset(flushTimerInterval)
		}
	}()
}

type exporter struct {
	file              *os.File
	bytesWritten      int
	bytesToFlushAfter int
	flushTimer        time.Timer
}

func (e *exporter) export(opCode OpCode, elapsedTime time.Duration, size uint32) {
	buf := []byte{opCode[0], opCode[1], 0, 0, 0, 0, 0, 0, 0, 0}
	binary.LittleEndian.PutUint32(buf[2:], uint32(elapsedTime))
	binary.LittleEndian.PutUint32(buf[6:], size)
	n, err := e.file.Write(buf)
	if err != nil {
		panic("could not write to benchmark file: " + err.Error())
	}

	e.bytesWritten += n
	if e.bytesWritten > e.bytesToFlushAfter {
		e.file.Sync()
		e.bytesWritten = 0
		e.flushTimer.Reset(flushTimerInterval)
	}
}

func (e *exporter) close() {
	e.file.Sync()
	e.file.Close()
}

func Finish() {
	fmt.Println("## StackSize: ", stackSize)

	fileWriter.close()
}
