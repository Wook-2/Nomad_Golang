package main

import (
	"bytes"
	"sync"
)

type SyncedBuffer struct {
	lock   sync.Mutex
	buffer bytes.Buffer
}

func main() {

	p := new(SyncedBuffer)	// type *SyncedBuffer
	var v SyncedBuffer	// type SyncedBuffer

}
