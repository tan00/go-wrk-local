package main

import (
	"encoding/hex"
	"fmt"
	"gmssl"
	"gmssl/sm3"
	"log"
	"sync"
)

type HashCtx struct {
	msg    []byte
	result []byte
}

type CipherCtx struct {
	key    []byte
	iv     []byte
	msg    []byte
	enc    bool //true 加密 flase 解密
	result []byte
}

var oncePrintSend, oncePrintRecv sync.Once
var (
	hashctx   HashCtx
	cipherCtx CipherCtx
)

func fsm3(ctx *HashCtx) {
	digest := sm3.New()
	ctx.result = digest.Sum(ctx.msg)
}

func fcipher(name string, ctx *CipherCtx) {
	keylen, err := gmssl.GetCipherKeyLength(name)
	if err != nil {
		log.Fatalf("gmssl.GetCipherKeyLength err %v\n", err)
	}

	ctx.key = ctx.key[0:keylen]
	gmsslCipher, err := gmssl.NewCipherContext(name, ctx.key, ctx.iv, ctx.enc)
	if err != nil {
		log.Fatalf("gmssl.NewCipherContext err %v \n", err)
	}

	buffer, err := gmsslCipher.Update(ctx.msg)
	ctx.result = append(ctx.result, buffer...)

	buffer, err = gmsslCipher.Final()
	if err != nil {
		log.Fatalf("gmssl.Final err %v\n", err)
	}
	ctx.result = append(ctx.result, buffer...)

}

//StartClient  url_: call function name
func StartClient(alg string, responseChan chan *Response, waitGroup *sync.WaitGroup, tc int) {
	defer waitGroup.Done()

	for {
		if len(responseChan) >= tc {
			break
		}

		timer := NewTimer()
		timer.Reset()

		switch alg { //AES-128-ECB|AES-256-ECB|SM3|SMS4-ECB
		case "SM3":
			ctx := hashctx
			fsm3(&ctx)
			oncePrintSend.Do(func() {
				if *printMsg {
					fmt.Println("==========================Print Msg==========================")
					log.Printf("alg: %s \n ", *algName)
					log.Printf("msg: %s \n ", hex.EncodeToString(ctx.msg))
					log.Printf("digest: %s \n ", hex.EncodeToString(ctx.result))
				}
			})
		case "AES-128-ECB", "AES-256-ECB", "SMS4-ECB":
			ctx := cipherCtx
			fcipher(alg, &ctx)
			oncePrintSend.Do(func() {
				if *printMsg {
					fmt.Println("==========================Print Msg==========================")
					log.Printf("alg: %s \n ", *algName)
					log.Printf("key: %s \n ", hex.EncodeToString(ctx.key))
					log.Printf("in msg: %s \n ", hex.EncodeToString(ctx.msg))
					log.Printf("out msg: %s \n ", hex.EncodeToString(ctx.result))
				}
			})
		default:
			log.Fatalf("alg  %s invalid\n", alg)
		}

		//解析返回
		respObj := &Response{Size: int64(*msgSize), StatusCode: 0, Error: false}
		t1 := timer.Duration()
		// if alg == "SMS4-ECB" || alg == "SMS4-CBC" {
		// 	t1 /= 2
		// }
		respObj.Duration = t1
		responseChan <- respObj
	}

}
