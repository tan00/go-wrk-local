package main

import (
	"flag"
	"os"
	"runtime"
)

type Config struct {
	Port  string
	Nodes []string
}

var (
	numThreads = flag.Int("t", 1, "the numbers of threads used")
	totalCalls = flag.Int("n", 1, "the total number of calls processed")
	printMsg   = flag.Bool("print", false, "print all in and out msg")
	msgSize    = flag.Int("size", -1, "msg size to calc Mbps")
	algName    = flag.String("alg", "", "the alg to run  [AES-128-ECB|AES-256-ECB|SM3|SMS4-ECB]")
)

func init() {
	flag.Parse()
	runtime.GOMAXPROCS(*numThreads) //一个gorotine 一个连接， 设置线程最大数量是为了限制同时运行的gorotine
}

func main() {
	if len(os.Args) == 1 { //If no argument specified
		flag.Usage() //Print usage
		os.Exit(1)
	}
	SingleNode(*algName)

}
