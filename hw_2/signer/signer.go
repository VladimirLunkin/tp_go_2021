package main

import (
	"fmt"
	"sort"
	"strconv"
	"sync"
	"time"
)

func ExecutePipeline(jobs ...job) {
	wg := &sync.WaitGroup{}
	in := make(chan interface{}, 1)
	for i, j := range jobs {
		wg.Add(1)
		out := make(chan interface{}, 1)
		go startWorker(j, in, out, wg)
		if i == 0 {
			in <- ""
		}
		in = out
	}
	fmt.Println(<-in)
	wg.Wait()
}

func startWorker(jobFunc job, in, out chan interface{}, waiter *sync.WaitGroup) {
	defer waiter.Done()
	defer close(out)

	jobFunc(in, out)
}

func SingleHash(in, out chan interface{}) {
	for ch := range in {
		data := strconv.Itoa(ch.(int))

		DataSignerOut1 := make(chan string)
		go DataSignerWorker(DataSignerCrc32, data, DataSignerOut1)

		DataSignerOut2 := make(chan string)
		go DataSignerWorker(DataSignerMd5, data, DataSignerOut2)
		s2 := <-DataSignerOut2

		DataSignerOut3 := make(chan string)
		go DataSignerWorker(DataSignerCrc32, s2, DataSignerOut3)

		s1 := <-DataSignerOut1
		s3 := <-DataSignerOut3

		out <- s1 + "~" + s3
	}
}

type DataSigner func(string2 string) string

func DataSignerWorker(dataSigner DataSigner, data string, out chan string) {
	out <- dataSigner(data)
}

const kMultiHash int = 6

func MultiHash(in, out chan interface{}) {
	for ch := range in {
		data := ch.(string)

		var chs [kMultiHash]chan string
		for i := range chs {
			chs[i] = make(chan string)
		}

		for th := 0; th < kMultiHash; th++ {
			go DataSignerWorker(DataSignerCrc32, strconv.Itoa(th)+data, chs[th])
		}

		var result string

		for i := range chs {
			result += <-chs[i]
		}

		out <- result
	}
}

func CombineResults(in, out chan interface{}) {
	var results []string
	for ch := range in {
		results = append(results, ch.(string))
	}

	sort.Strings(results)

	var res string

	for _, str := range results {
		res += str + "_"
	}

	out <- res[:len(res)-1]
}

func qwe(in, out chan interface{}) {
	out <- 0
	out <- 1
}

func main() {
	tStart := time.Now()
	ExecutePipeline(qwe, SingleHash, MultiHash, CombineResults)
	tFinish := time.Now()
	fmt.Println(tFinish.UnixMicro() - tStart.UnixMicro())
}
