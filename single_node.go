package main

import (
	"bytes"
	"sync"
)

func SingleNode(toCall string) []byte {
	var (
		responseChannel = make(chan *Response, *totalCalls*2)
	)

	ffbyte := make([]byte, 1)
	ffbyte[0] = 0xff
	hashctx.msg = bytes.Repeat(ffbyte, *msgSize)
	cipherCtx.key = bytes.Repeat(ffbyte, 32)
	cipherCtx.iv = bytes.Repeat(ffbyte, 32)
	cipherCtx.msg = bytes.Repeat(ffbyte, *msgSize)
	cipherCtx.enc = true

	benchTime := NewTimer()
	benchTime.Reset()
	//TODO check ulimit
	wg := &sync.WaitGroup{}

	for i := 0; i < *numThreads*2; i++ {
		go StartClient(
			toCall,
			responseChannel,
			wg,
			*totalCalls,
		)
		wg.Add(1)
	}

	wg.Wait()

	result := CalcStats(
		responseChannel,
		benchTime.Duration(),
	)
	return result
}
