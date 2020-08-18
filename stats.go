package main

import (
	"encoding/json"
	"fmt"
	"sort"
)

type Stats struct {
	Alg string
	//Connections int
	Threads     int
	AvgDuration float64
	Duration    float64
	Sum         float64
	Times       []int
	Transferred int64
	Success     int64
	Errors      int64
}

func CalcStats(responseChannel chan *Response, duration int64) []byte {

	stats := &Stats{
		Alg: *algName,

		Threads:     *numThreads,
		Times:       make([]int, len(responseChannel)),
		Duration:    float64(duration),
		AvgDuration: float64(duration),
	}

	i := 0
	for res := range responseChannel {
		switch {
		case res.StatusCode == 0:
			stats.Success++
		default:
			stats.Errors++
		}

		stats.Sum += float64(res.Duration)
		stats.Times[i] = int(res.Duration)
		i++

		stats.Transferred += res.Size

		if res.Error {
			stats.Errors++
		}

		if len(responseChannel) == 0 {
			break
		}
	}

	sort.Ints(stats.Times)

	PrintStats(stats)
	b, err := json.Marshal(&stats)
	if err != nil {
		fmt.Println(err)
	}
	return b
}

func PrintStats(allStats *Stats) {
	sort.Ints(allStats.Times)
	total := float64(len(allStats.Times))
	totalInt := int64(total)
	fmt.Println("==========================BENCHMARK==========================")
	fmt.Printf("Alg:\t\t\t\t%s\n\n", allStats.Alg)
	fmt.Printf("Used Threads:\t\t\t%d\n", allStats.Threads)
	fmt.Printf("Total number of calls:\t\t%d\n\n", totalInt)
	fmt.Println("===========================TIMINGS===========================")
	fmt.Printf("Total time passed:\t\t%.2fs\n", allStats.AvgDuration/1E6)
	fmt.Printf("Avg time per request:\t\t%.2fus\n", allStats.Sum/total)
	fmt.Printf("Requests per second:\t\t%.2f\n", total/(allStats.AvgDuration/1E6))
	fmt.Printf("Median time per request:\t%.2fus\n", float64(allStats.Times[(totalInt-1)/2]))
	fmt.Printf("99th percentile time:\t\t%.2fus\n", float64(allStats.Times[(totalInt/100*99)]))
	fmt.Printf("Slowest time for request:\t%.2fus\n\n", float64(allStats.Times[totalInt-1]))
	fmt.Println("=============================DATA=============================")
	fmt.Printf("Total sizes:\t\t%d\n", allStats.Transferred)
	fmt.Printf("Avg response body per request:\t\t%.2f Byte\n", float64(allStats.Transferred)/total)
	tr := float64(allStats.Transferred) / (allStats.AvgDuration / 1E6)
	fmt.Printf("Transfer rate per second:\t\t%.2f Mb/s\n", 8*tr/1E6)
	fmt.Println("==========================RESPONSES==========================")
	fmt.Printf("Success Responses:\t\t%d\t(%.2f%%)\n", allStats.Success, float64(allStats.Success)/total*1e2)
	fmt.Printf("Error Responses:\t\t%d\t(%.2f%%)\n\n\n", allStats.Errors, float64(allStats.Errors)/total*1e2)
}
